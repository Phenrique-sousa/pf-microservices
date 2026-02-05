package shipping

import (
	"context"
	"fmt"
	"time"

	shippingpb "github.com/phenrique-sousa/pf-microservices-proto/golang/shipping"
	"github.com/phenrique-sousa/pf-microservices/order/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	client shippingpb.ShippingClient
}

func NewAdapter(url string) (*Adapter, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("shipping connection error: %v", err)
	}
	return &Adapter{client: shippingpb.NewShippingClient(conn)}, nil
}

func (a Adapter) CreateShipping(orderId int64, items []ports.ShippingItem) (int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var pbItems []*shippingpb.ShippingItem
	for _, it := range items {
		pbItems = append(pbItems, &shippingpb.ShippingItem{
			ProductCode: it.ProductCode,
			Quantity:    it.Quantity,
		})
	}

	resp, err := a.client.Create(ctx, &shippingpb.CreateShippingRequest{
		OrderId: orderId,
		Items:   pbItems,
	})
	if err != nil {
		return 0, err
	}
	return resp.DeliveryDays, nil
}
