package grpc

import (
	"context"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"server/src/domain"
	pb "server/src/transport/pb"
	"time"
)

func encodeResponse(_ context.Context, resp domain.RemotePayment) (proto.Message, error) {
	response := &pb.Payment{
		Type:             resp.Type,
		Status:           resp.Status,
		AccountNumber:    resp.AccountNumber,
		Amount:           resp.Amount.String(),
		PayerShortName:   resp.PayerShortName,
		PayerPhoneNumber: resp.PayerPhoneNumber,
		ReceiptNumber:    resp.ReceiptNumber,
		Description:      resp.Description,
		CreatedAt:        timestamppb.New(time.Time(resp.CreatedAt)),
	}

	return response, nil
}
