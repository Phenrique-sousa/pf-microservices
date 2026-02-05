package api

import (
	"github.com/phenrique-sousa/pf-microservices/shipping/internal/application/core/domain"
	"github.com/phenrique-sousa/pf-microservices/shipping/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a Application) CreateShipping(orderId int64, items []domain.ShippingItem) (domain.Shipping, error) {
	// regra simples: 3 dias padrão; 7 dias se tiver 3+ itens
	days := int32(3)
	if len(items) >= 3 {
		days = 7
	}

	s := domain.NewShipping(orderId, items, days)
	if err := a.db.Save(&s); err != nil {
		return domain.Shipping{}, err
	}
	return s, nil
}
