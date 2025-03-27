package transport

import (
	"github.com/google/uuid"
	"server/src/domain"
)

type (
	ListenPaymentChangesRequest struct {
		QrID uuid.UUID `json:"qrId"`
	}
	ListenPaymentChangesResponse struct {
		Channel <-chan domain.RemotePayment `json:"channel"`
	}
)
