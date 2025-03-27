package domain

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type MyTime time.Time

func (t *MyTime) UnmarshalJSON(data []byte) error {
	// Define the format for unmarshalling
	layout := `"2006-01-02T15:04:05.999999"`
	// Parse the time
	parsedTime, err := time.Parse(layout, string(data))
	if err != nil {
		return err
	}
	*t = MyTime(parsedTime)
	return nil
}

type RemotePayment struct {
	ID                 string          `json:"id"`
	Type               string          `json:"type"`        // qr or invoice
	ServiceMappingType string          `json:"serviceType"` // selling_point or company_service
	ServiceMappingID   uuid.UUID       `json:"serviceId"`   // selling_point_id or company_service_id
	CompanyID          uuid.UUID       `json:"companyId"`
	ManagerID          *uuid.UUID      `json:"managerId"`
	AccountNumber      string          `json:"accountNumber"`
	Amount             decimal.Decimal `json:"amount"` // payment amount
	ReceiptNumber      *string         `json:"receiptNumber"`
	PayerShortName     string          `json:"payerShortName"`
	PayerPhoneNumber   string          `json:"payerPhoneNumber"`
	Description        *string         `json:"description"`
	Status             string          `json:"status"`
	TransactionID      int64           `json:"transactionId"`
	QrID               *uuid.UUID      `json:"qrId"`
	CreatedAt          MyTime          `json:"createdAt"`
}
