package api

import (
	"github.com/phenrique-sousa/pf-microservices/order/internal/application/core/domain"
	"github.com/phenrique-sousa/pf-microservices/order/internal/ports"
)

type Application struct {
	db       ports.DBPort
	payment  ports.PaymentPort
	shipping ports.ShippingPort
}

func NewApplication(
	db ports.DBPort,
	payment ports.PaymentPort,
	shipping ports.ShippingPort,
) *Application {
	return &Application{
		db:       db,
		payment:  payment,
		shipping: shipping,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	// 1 salva o pedido
	if err := a.db.Save(&order); err != nil {
		return domain.Order{}, err
	}

	// 2 calcula o valor total do pedido
	var total float32 = 0
	for _, it := range order.OrderItems {
		total += it.UnitPrice * float32(it.Quantity)
	}

	// 3 chama o serviço de Payment
	_, err := a.payment.CreatePayment(order.CustomerID, order.ID, total)
	if err != nil {
		return domain.Order{}, err
	}

	// 4 prepara itens para o Shipping
	var shippingItems []ports.ShippingItem
	for _, it := range order.OrderItems {
		shippingItems = append(shippingItems, ports.ShippingItem{
			ProductCode: it.ProductCode,
			Quantity:    int32(it.Quantity),
		})
	}

	// 5 chama o serviço de Shipping
	_, err = a.shipping.CreateShipping(order.ID, shippingItems)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}
