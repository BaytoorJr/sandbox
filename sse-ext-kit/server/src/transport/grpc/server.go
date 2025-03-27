package grpc

import (
	kitgrpc "github.com/einouqo/ext-kit/transport/grpc"
	gokittransport "github.com/go-kit/kit/transport"
	"github.com/go-kit/log"
	"server/src/middleware"
	pb "server/src/transport/pb"
)

type ServiceServer struct {
	listenPaymentChanges kitgrpc.HandlerInnerStream

	pb.UnimplementedPaymentServiceServer
}

func NewGRPCServer(svc *middleware.Endpoint, logger log.Logger) *ServiceServer {
	var opts = []kitgrpc.ServerOption{
		kitgrpc.WithServerErrorHandler(gokittransport.NewLogErrorHandler(logger)),
	}

	return &ServiceServer{
		listenPaymentChanges: kitgrpc.NewServerInnerStream(
			svc.ListenPaymentChanges,
			decodeRequest,
			encodeResponse,
			opts...,
		),
	}
}

func (ss *ServiceServer) ListenPaymentChanges(req *pb.ListenPaymentChangesRequest, stream pb.PaymentService_ListenPaymentChangesServer) error {
	_, err := ss.listenPaymentChanges.ServeInnerStream(req, stream)
	if err != nil {
		return err
	}

	return nil
}
