package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	orderpb "github.com/phenrique-sousa/pf-microservices-proto/golang/order"
	"github.com/phenrique-sousa/pf-microservices/order/internal/application/core/api"
	"github.com/phenrique-sousa/pf-microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  *api.Application
	port int
	orderpb.UnimplementedOrderServer
}

func NewAdapter(api *api.Application, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Create(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	var items []domain.OrderItem
	for _, it := range req.OrderItems {
		items = append(items, domain.OrderItem{
			ProductCode: it.ProductCode,
			UnitPrice:   it.UnitPrice,
			Quantity:    it.Quantity,
		})
	}

	order := domain.NewOrder(int64(req.CostumerId), items)
	result, err := a.api.PlaceOrder(order)

	if err != nil {
		return nil, err
	}

	return &orderpb.CreateOrderResponse{OrderId: int32(result.ID)}, nil

}

func (a Adapter) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	orderpb.RegisterOrderServer(s, a)
	reflection.Register(s)

	log.Printf("Order gRPC running on port %d", a.port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
