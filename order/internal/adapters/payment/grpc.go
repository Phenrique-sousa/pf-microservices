package payment

import (
	"context"
	"fmt"
	"time"

	paymentpb "github.com/phenrique-sousa/pf-microservices-proto/golang/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	client paymentpb.PaymentClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	conn, err := grpc.NewClient(paymentServiceUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("payment connection error: %v", err)
	}

	client := paymentpb.NewPaymentClient(conn)
	return &Adapter{client: client}, nil
}

func (a Adapter) CreatePayment(userId, orderId int64, total float32) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := a.client.Create(ctx, &paymentpb.CreatePaymentRequest{
		UserId:     userId,
		OrderId:    orderId,
		TotalPrice: total,
	})
	if err != nil {
		return 0, err
	}
	return resp.PaymentId, nil
}
