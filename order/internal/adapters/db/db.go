package db

import (
	"fmt"

	"github.com/phenrique-sousa/pf-microservices/order/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderID     uint
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dsn string) (*Adapter, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connection error: %v", err)
	}

	if err := db.AutoMigrate(&Order{}, &OrderItem{}); err != nil {
		return nil, fmt.Errorf("migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (a Adapter) Save(order *domain.Order) error {
	var items []OrderItem
	for _, it := range order.OrderItems {
		items = append(items, OrderItem{
			ProductCode: it.ProductCode,
			UnitPrice:   it.UnitPrice,
			Quantity:    it.Quantity,
		})
	}

	model := Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: items,
	}

	res := a.db.Create(&model)
	if res.Error == nil {
		order.ID = int64(model.ID)
	}
	return res.Error
}
