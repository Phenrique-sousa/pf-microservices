package main

import (
	"log"

	"github.com/phenrique-sousa/pf-microservices/payment/config"
	dbadapter "github.com/phenrique-sousa/pf-microservices/payment/internal/adapters/db"
	grpcadapter "github.com/phenrique-sousa/pf-microservices/payment/internal/adapters/grpc"
	"github.com/phenrique-sousa/pf-microservices/payment/internal/application/core/api"
)

func main() {
	db, err := dbadapter.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}

	app := api.NewApplication(db)
	server := grpcadapter.NewAdapter(app, 3001)
	server.Run()
}
