package ports

import "github.com/phenrique-sousa/pf-microservices/payment/internal/application/core/domain"

type DBPort interface {
	Save(payment *domain.Payment) error
}
