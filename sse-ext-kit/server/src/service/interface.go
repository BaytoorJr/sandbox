package service

import (
	"context"
	"github.com/einouqo/ext-kit/endpoint"
	"server/src/domain"
	"server/src/transport"
)

type Service interface {
	ListenPaymentChanges(ctx context.Context, req *transport.ListenPaymentChangesRequest) (endpoint.Receive[domain.RemotePayment], error)
}
