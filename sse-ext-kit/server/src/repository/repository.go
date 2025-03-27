package repository

import (
	"context"
	"github.com/google/uuid"
	"server/src/domain"
)

type RemotePaymentRepository interface {
	ListenPaymentChanges(ctx context.Context, qrID uuid.UUID, paymentChan chan<- domain.RemotePayment) error
}
