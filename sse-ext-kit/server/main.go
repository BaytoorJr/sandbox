package main

import (
	"context"
	"fmt"
	"github.com/go-kit/log/level"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"server/src/middleware"
	"server/src/repository/postgres"
	"server/src/service"
	paymentGRPC "server/src/transport/grpc"
	pb "server/src/transport/pb"
	"sync/atomic"
	"syscall"

	"google.golang.org/grpc"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	grpcLocale "gitlab.globerce.com/freedom-business/libs/shared-libs/middleware/locale/grpc"
	liblogger "gitlab.globerce.com/freedom-business/libs/shared-libs/utils/logger"

	"server/src/config"
)

func main() {
	// main ctx
	ctx := context.Background()

	// init structured logger for the service
	logger := liblogger.NewServiceLogger("sse-api")
	_ = level.Info(logger).Log("msg", "service started")

	cfg, err := config.InitConfigs()
	if err != nil {
		_ = level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	mainStore, err := postgres.New(ctx, cfg.PostgresConfig, logger)
	if err != nil {
		_ = level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	svc := service.NewService(mainStore, logger)
	svc = middleware.NewLoggingMiddleware(logger)(svc)
	svc = middleware.NewInstrumentingMiddleware(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "sse_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, []string{"method"}),
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "sse_service",
			Name:      "error_count",
			Help:      "Number of error requests received.",
		}, []string{"method"}),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "sse_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, []string{"method"}),
	)(svc)

	srvDone := atomic.Value{}
	srvDone.Store(false)

	svcEndpoints := middleware.MakeEndpoints(svc)

	grpcServer := paymentGRPC.NewGRPCServer(svcEndpoints, nil)
	baseGRPCServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcLocale.FetchLanguageUnaryInterceptor,
		),
	)
	pb.RegisterPaymentServiceServer(baseGRPCServer, grpcServer)
	reflection.Register(baseGRPCServer)

	grpcListener, err := net.Listen("tcp", cfg.GRPCConfig.ListenAddr)
	if err != nil {
		_ = level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	errs := make(chan error)
	// make chan for syscall
	go func() {
		c := make(chan os.Signal, config.ChannelDefaultSize)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// init gRPC server
	go func() {
		_ = level.Info(logger).Log("transport", "gRPC", "port", cfg.GRPCConfig.ListenAddr)
		errs <- baseGRPCServer.Serve(grpcListener)
	}()

	defer func() {
		_ = level.Info(logger).Log("msg", "service ended")
	}()

	_ = level.Error(logger).Log("exit", <-errs)

	baseGRPCServer.GracefulStop()
	_ = grpcListener.Close()
}
