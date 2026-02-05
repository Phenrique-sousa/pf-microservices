package ports

type ShippingItem struct {
	ProductCode string
	Quantity    int32
}

type ShippingPort interface {
	CreateShipping(orderId int64, items []ShippingItem) (int32, error)
}
