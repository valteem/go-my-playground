package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"

	"webapi/product-catalog/app"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}
}

func main() {
	app.Run(context.Background())
}
