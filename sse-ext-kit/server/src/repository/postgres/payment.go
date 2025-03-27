package postgres

import (
	"context"
	"encoding/json"
	"github.com/go-kit/log"
	"github.com/google/uuid"
	"github.com/jackc/pgproto3/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.globerce.com/freedom-business/libs/shared-libs/errors"
	"server/src/domain"
)

type RemotePayment struct {
	db     *pgxpool.Pool
	logger log.Logger
}

func (r *RemotePayment) ListenPaymentChanges(ctx context.Context,
	qrID uuid.UUID, paymentChan chan<- domain.RemotePayment) error {
	conn, err := r.db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "LISTEN payment_row_change_channel")
	if err != nil {
		return errors.InternalServerError.SetDevMessage(err.Error()).AddClientError(ctx, "TECH_500", nil)
	}

	pgConn := conn.Conn().PgConn()

	for {
		select {
		case <-ctx.Done():
			_ = pgConn.Close(ctx)
			conn.Release()
			return nil
		default:
			msg, err := pgConn.ReceiveMessage(ctx)
			if err != nil {
				return errors.DBReadError.SetDevMessage(err.Error()).AddClientError(ctx, "TECH_500", nil)
			}

			notification, ok := msg.(*pgproto3.NotificationResponse)
			if !ok {
				return errors.InternalServerError.SetDevMessage("received unexpected notification message")
			}

			var payment domain.RemotePayment
			err = json.Unmarshal([]byte(notification.Payload), &payment)
			if err != nil {
				return errors.SerializeError.SetDevMessage(err.Error()).AddClientError(ctx, "TECH_500", nil)
			}

			if payment.QrID != nil && *payment.QrID == qrID {
				select {
				case <-ctx.Done():
					return nil
				case paymentChan <- payment:
				}
			}
		}
	}
}
