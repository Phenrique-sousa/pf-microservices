package config

import (
	"log"
	"os"
)

func GetDataSourceURL() string {
	dsn := os.Getenv("DATA_SOURCE_URL")
	if dsn == "" {
		log.Fatal("DATA_SOURCE_URL environment variable is missing")
	}
	return dsn
}
