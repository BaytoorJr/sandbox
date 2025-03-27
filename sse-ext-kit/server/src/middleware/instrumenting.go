package middleware

import (
	"context"
	"github.com/einouqo/ext-kit/endpoint"
	"github.com/go-kit/kit/metrics"
	"server/src/domain"
	"server/src/service"
	"server/src/transport"
	"time"
)

type instrumentingMiddleware struct {
	next           service.Service
	requestCount   metrics.Counter
	requestError   metrics.Counter
	requestLatency metrics.Histogram
}

func NewInstrumentingMiddleware(counter, counterErr metrics.Counter, latency metrics.Histogram) Middleware {
	return func(next service.Service) service.Service {
		return &instrumentingMiddleware{
			next:           next,
			requestCount:   counter,
			requestError:   counterErr,
			requestLatency: latency,
		}
	}
}

func (m *instrumentingMiddleware) writeMethodMetrics(begin time.Time, method string, err error) {
	m.requestCount.With("method", method).Add(1)
	m.requestLatency.With("method", method).Observe(time.Since(begin).Seconds())

	if err != nil {
		m.requestError.With("method", method).Add(1)
	}
}

func (m *instrumentingMiddleware) ListenPaymentChanges(ctx context.Context, req *transport.ListenPaymentChangesRequest) (_ endpoint.Receive[domain.RemotePayment], err error) {
	defer func(now time.Time) {
		m.writeMethodMetrics(now, "listenPaymentChanges", err)
	}(time.Now())

	return m.next.ListenPaymentChanges(ctx, req)
}
