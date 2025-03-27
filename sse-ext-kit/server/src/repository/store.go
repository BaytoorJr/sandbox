package repository

type Storage interface {
	RemotePayment() RemotePaymentRepository
}
