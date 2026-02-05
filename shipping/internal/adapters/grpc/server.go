package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	shippingpb "github.com/phenrique-sousa/pf-microservices-proto/golang/shipping"
	"github.com/phenrique-sousa/pf-microservices/shipping/internal/application/core/api"
	"github.com/phenrique-sousa/pf-microservices/shipping/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  *api.Application
	port int
	shippingpb.UnimplementedShippingServer
}

func NewAdapter(api *api.Application, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Create(ctx context.Context, req *shippingpb.CreateShippingRequest) (*shippingpb.CreateShippingResponse, error) {
	var items []domain.ShippingItem
	for _, it := range req.Items {
		items = append(items, domain.ShippingItem{
			ProductCode: it.ProductCode,
			Quantity:    it.Quantity,
		})
	}

	result, err := a.api.CreateShipping(req.OrderId, items)
	if err != nil {
		return nil, err
	}

	return &shippingpb.CreateShippingResponse{DeliveryDays: result.Days}, nil
}

func (a Adapter) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	shippingpb.RegisterShippingServer(s, a)
	reflection.Register(s)

	log.Printf("Shipping gRPC running on port %d", a.port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
