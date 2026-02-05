package main

import (
	"log"

	"github.com/phenrique-sousa/pf-microservices/order/config"
	dbadapter "github.com/phenrique-sousa/pf-microservices/order/internal/adapters/db"
	grpcadapter "github.com/phenrique-sousa/pf-microservices/order/internal/adapters/grpc"
	paymentadapter "github.com/phenrique-sousa/pf-microservices/order/internal/adapters/payment"
	shippingadapter "github.com/phenrique-sousa/pf-microservices/order/internal/adapters/shipping"
	"github.com/phenrique-sousa/pf-microservices/order/internal/application/core/api"
)

func main() {
	// Database adapter
	db, err := dbadapter.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}

	// Payment adapter
	payment, err := paymentadapter.NewAdapter(config.GetPaymentServiceURL())
	if err != nil {
		log.Fatalf("Payment error: %v", err)
	}

	// Shipping adapter
	shipping, err := shippingadapter.NewAdapter(config.GetShippingServiceURL())
	if err != nil {
		log.Fatalf("Shipping error: %v", err)
	}

	// Application core
	app := api.NewApplication(db, payment, shipping)

	// gRPC server
	server := grpcadapter.NewAdapter(app, 3000)
	server.Run()
}
