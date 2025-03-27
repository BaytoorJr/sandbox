package grpc

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"server/src/transport"
	pb "server/src/transport/pb"
)

func decodeRequest(_ context.Context, request proto.Message) (*transport.ListenPaymentChangesRequest, error) {
	req := request.(*pb.ListenPaymentChangesRequest)

	qrId, _ := uuid.Parse(req.GetQrID())

	r := transport.ListenPaymentChangesRequest{
		QrID: qrId,
	}

	return &r, nil
}
