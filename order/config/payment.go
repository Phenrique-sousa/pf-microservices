package config

import (
	"log"
	"os"
)

func GetPaymentServiceURL() string {
	v := os.Getenv("PAYMENT_SERVICE_URL")
	if v == "" {
		log.Fatal("PAYMENT_SERVICE_URL environment variable is missing")
	}
	return v
}
