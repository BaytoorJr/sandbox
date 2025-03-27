package middleware

import (
	"context"
	"github.com/einouqo/ext-kit/endpoint"
	"github.com/go-kit/log"
	"server/src/domain"
	"server/src/service"
	"server/src/transport"
	"time"
)

type loggingMiddleware struct {
	next   service.Service
	logger log.Logger
}

func (l *loggingMiddleware) logMethod(begin time.Time, method string, err error) {
	_ = l.logger.Log("method", method, "took", time.Since(begin), "err", err)
}

func NewLoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.Service) service.Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

func (l *loggingMiddleware) ListenPaymentChanges(ctx context.Context, req *transport.ListenPaymentChangesRequest) (_ endpoint.Receive[domain.RemotePayment], err error) {
	now := time.Now()
	defer func(now time.Time) {
		l.logMethod(now, "listenPaymentChanges", err)
	}(now)

	return l.next.ListenPaymentChanges(ctx, req)
}
