package service

import (
	"context"
	"github.com/einouqo/ext-kit/endpoint"
	"github.com/go-kit/log/level"
	"server/src/domain"
	"server/src/transport"
)

func (s *service) ListenPaymentChanges(ctx context.Context, req *transport.ListenPaymentChangesRequest) (endpoint.Receive[domain.RemotePayment], error) {
	paymentChan := make(chan domain.RemotePayment, 10)
	errChan := make(chan error, 1)
	go func() {
		select {
		case <-ctx.Done():
			return
		default:
			err := s.store.RemotePayment().ListenPaymentChanges(ctx, req.QrID, paymentChan)
			if err != nil {
				errChan <- err
				_ = level.Error(s.logger).Log("listen_payment_changes_error", err)
				return
			}
		}
	}()

	return func() (domain.RemotePayment, error) {
		resp, ok := <-paymentChan
		if ok {
			return resp, nil
		}

		err, ok := <-errChan
		if ok {
			return domain.RemotePayment{}, err
		}

		return domain.RemotePayment{}, endpoint.StreamDone
	}, nil
}
