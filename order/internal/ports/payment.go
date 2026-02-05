package ports

type PaymentPort interface {
	CreatePayment(userId, orderId int64, total float32) (int64, error)
}
