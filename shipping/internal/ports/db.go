package ports

import "github.com/phenrique-sousa/pf-microservices/shipping/internal/application/core/domain"

type DBPort interface {
	Save(shipping *domain.Shipping) error
}
