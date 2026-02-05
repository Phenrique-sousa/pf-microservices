package db

import (
	"fmt"

	"github.com/phenrique-sousa/pf-microservices/payment/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	UserID  int64
	OrderID int64
	Total   float32
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dsn string) (*Adapter, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connection error: %v", err)
	}

	if err := db.AutoMigrate(&Payment{}); err != nil {
		return nil, fmt.Errorf("migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (a Adapter) Save(payment *domain.Payment) error {
	model := Payment{
		UserID:  payment.UserID,
		OrderID: payment.OrderID,
		Total:   payment.Total,
	}

	res := a.db.Create(&model)
	if res.Error == nil {
		payment.ID = int64(model.ID)
	}
	return res.Error
}
