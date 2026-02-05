package domain

import "time"

type ShippingItem struct {
	ProductCode string
	Quantity    int32
}

type Shipping struct {
	ID        int64
	OrderID   int64
	Days      int32
	CreatedAt int64
	Items     []ShippingItem
}

func NewShipping(orderId int64, items []ShippingItem, days int32) Shipping {
	return Shipping{
		OrderID:   orderId,
		Days:      days,
		Items:     items,
		CreatedAt: time.Now().Unix(),
	}
}
