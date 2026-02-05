package domain

import "time"

type Payment struct {
	ID        int64
	UserID    int64
	OrderID   int64
	Total     float32
	CreatedAt int64
}

func NewPayment(userId, orderId int64, total float32) Payment {
	return Payment{
		UserID:    userId,
		OrderID:   orderId,
		Total:     total,
		CreatedAt: time.Now().Unix(),
	}
}
