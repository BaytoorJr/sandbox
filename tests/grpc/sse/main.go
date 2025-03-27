package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgproto3/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	pb "golangProject.com/grpc/proto/events"
	"golangProject.com/grpc/sse/domain"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

const (
	user     = "yelnur"
	pass     = "BTR2002"
	host     = "localhost"
	port     = "5432"
	dbName   = "globerce"
	maxConns = 100
)

type eventServiceServer struct {
	pb.UnimplementedEventServiceServer
}

func (s *eventServiceServer) StreamEvents(req *pb.Empty, stream pb.EventService_StreamEventsServer) error {
	ctx := stream.Context()
	paymentChan := make(chan domain.SseTest, 10)
	go func() {
		err := WatchChanges(ctx, paymentChan, 1)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	for payment := range paymentChan {
		log.Println("received from chan: ", payment.ID, payment.Status)

		if err := stream.Send(&pb.PaymentStatus{
			Id:     int32(payment.ID),
			Status: payment.Status,
			Amount: "",
		}); err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEventServiceServer(grpcServer, &eventServiceServer{})

	fmt.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func WatchChanges(ctx context.Context, paymentChan chan<- domain.SseTest, rowID int) error {
	defer close(paymentChan)
	pool, err := connectToDB(ctx)
	if err != nil {
		return err
	}
	defer pool.Close()

	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "LISTEN payment_row_change_channel")
	if err != nil {
		return err
	}

	log.Println("Started listening for database changes...")

	pgConn := conn.Conn().PgConn()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			err = conn.Ping(ctx)
			if err != nil {
				log.Println(err)
			}

			msg, err := pgConn.ReceiveMessage(ctx)
			if err != nil {
				log.Println("Error receiving notification:", err)
				return err
			}

			notification, ok := msg.(*pgproto3.NotificationResponse)
			if !ok {
				return errors.New("received unexpected notification message")
			}

			log.Printf("Received notification: %t", notification.Payload)

			var pmnt domain.SseTest

			err = json.Unmarshal([]byte(notification.Payload), &pmnt)
			if err != nil {
				log.Println(err)
				return err
			}

			log.Println(pmnt.ID, pmnt.Status)

			if rowID == pmnt.ID {
				paymentChan <- pmnt
				log.Println("changes written")
			}
		}
	}
}

func parseNotification(payload []byte) (domain.SseTest, error) {
	// Assume JSON payload. Adapt parsing as per your use case.
	var notification domain.SseTest
	err := json.Unmarshal(payload, &notification)
	if err != nil {
		return domain.SseTest{}, err
	}
	return notification, nil
}

func connectToDB(ctx context.Context) (*pgxpool.Pool, error) {
	dsn := "postgres://" +
		user + ":" +
		pass + "@" +
		host + ":" +
		port + "/" +
		dbName +
		"?sslmode=disable" +
		"&pool_max_conns=" + strconv.Itoa(maxConns) +
		"&pool_max_conn_lifetime=900s" +
		"&pool_max_conn_lifetime_jitter=300s" +
		"&pool_max_conn_idle_time=600s"

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
