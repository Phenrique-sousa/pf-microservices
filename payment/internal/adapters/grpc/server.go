package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	paymentpb "github.com/phenrique-sousa/pf-microservices-proto/golang/payment"
	"github.com/phenrique-sousa/pf-microservices/payment/internal/application/core/api"
	"github.com/phenrique-sousa/pf-microservices/payment/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	api  *api.Application
	port int
	paymentpb.UnimplementedPaymentServer
}

func NewAdapter(api *api.Application, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Create(ctx context.Context, req *paymentpb.CreatePaymentRequest) (*paymentpb.CreatePaymentResponse, error) {
	p := domain.NewPayment(req.UserId, req.OrderId, req.TotalPrice)

	result, err := a.api.CreatePayment(p)
	if err != nil {
		// Regra: se total > 1000, retornar INVALID_ARGUMENT
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &paymentpb.CreatePaymentResponse{PaymentId: result.ID}, nil
}

func (a Adapter) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	paymentpb.RegisterPaymentServer(s, a)
	reflection.Register(s)

	log.Printf("Payment gRPC running on port %d", a.port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
