package db

import (
	"fmt"

	"github.com/phenrique-sousa/pf-microservices/shipping/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Shipping struct {
	gorm.Model
	OrderID int64
	Days    int32
	Items   []ShippingItem
}

type ShippingItem struct {
	gorm.Model
	ProductCode string
	Quantity    int32
	ShippingID  uint
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dsn string) (*Adapter, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connection error: %v", err)
	}

	if err := db.AutoMigrate(&Shipping{}, &ShippingItem{}); err != nil {
		return nil, fmt.Errorf("migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (a Adapter) Save(s *domain.Shipping) error {
	var items []ShippingItem
	for _, it := range s.Items {
		items = append(items, ShippingItem{
			ProductCode: it.ProductCode,
			Quantity:    it.Quantity,
		})
	}

	model := Shipping{
		OrderID: s.OrderID,
		Days:    s.Days,
		Items:   items,
	}

	res := a.db.Create(&model)
	if res.Error == nil {
		s.ID = int64(model.ID)
	}
	return res.Error
}
