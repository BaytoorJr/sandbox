package middleware

import (
	"context"
	"github.com/einouqo/ext-kit/endpoint"
	"server/src/domain"
	"server/src/service"
	"server/src/transport"
)

type Endpoint struct {
	ListenPaymentChanges endpoint.InnerStream[*transport.ListenPaymentChangesRequest, domain.RemotePayment]
}

func MakeEndpoints(s service.Service) *Endpoint {
	return &Endpoint{
		ListenPaymentChanges: makeListenPaymentChanges(s),
	}
}

func makeListenPaymentChanges(s service.Service) endpoint.InnerStream[*transport.ListenPaymentChangesRequest, domain.RemotePayment] {
	return func(ctx context.Context, request *transport.ListenPaymentChangesRequest) (endpoint.Receive[domain.RemotePayment], error) {
		return s.ListenPaymentChanges(ctx, request)
	}
}
