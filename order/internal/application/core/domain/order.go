package domain

import "time"

type OrderItem struct {
	ProductCode string
	UnitPrice   float32
	Quantity    int32
}

type Order struct {
	ID         int64
	CustomerID int64
	Status     string
	OrderItems []OrderItem
	CreatedAt  int64
}

func NewOrder(customerId int64, items []OrderItem) Order {
	return Order{
		CustomerID: customerId,
		Status:     "Pending",
		OrderItems: items,
		CreatedAt:  time.Now().Unix(),
	}
}
