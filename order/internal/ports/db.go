package ports

import "github.com/phenrique-sousa/pf-microservices/order/internal/application/core/domain"

type DBPort interface {
	Save(order *domain.Order) error
}
