package config

import (
	"log"
	"os"
)

func GetShippingServiceURL() string {
	v := os.Getenv("SHIPPING_SERVICE_URL")
	if v == "" {
		log.Fatal("SHIPPING_SERVICE_URL environment variable is missing")
	}
	return v
}
