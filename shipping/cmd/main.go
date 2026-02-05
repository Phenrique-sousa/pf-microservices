package main

import (
	"log"

	"github.com/phenrique-sousa/pf-microservices/shipping/config"
	dbadapter "github.com/phenrique-sousa/pf-microservices/shipping/internal/adapters/db"
	grpcadapter "github.com/phenrique-sousa/pf-microservices/shipping/internal/adapters/grpc"
	"github.com/phenrique-sousa/pf-microservices/shipping/internal/application/core/api"
)

func main() {
	db, err := dbadapter.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}

	app := api.NewApplication(db)
	server := grpcadapter.NewAdapter(app, 3002)
	server.Run()
}
