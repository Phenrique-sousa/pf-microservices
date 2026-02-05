package api

import (
	"fmt"

	"github.com/phenrique-sousa/pf-microservices/payment/internal/application/core/domain"
	"github.com/phenrique-sousa/pf-microservices/payment/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a Application) CreatePayment(p domain.Payment) (domain.Payment, error) {
	// Regra: não permitir pagamento acima de 1000
	if p.Total > 1000 {
		return domain.Payment{}, fmt.Errorf("payment total exceeds 1000")
	}

	if err := a.db.Save(&p); err != nil {
		return domain.Payment{}, err
	}
	return p, nil
}
