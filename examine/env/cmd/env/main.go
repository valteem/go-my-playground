package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	envs = []string{"APP_VERSION", "DB_MAX_POOL_SIZE", "DB_URL", "HTTP_PORT", "HTTP_READ_TIMEOUT", "HTTP_WRITE_TIMEOUT"}
)

func main() {

	dir, _ := os.Getwd()
	log.Printf("%s", dir)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load env variables: %v", err)
	}

	for _, env := range envs {
		log.Printf("%s = %s", env, os.Getenv(env))
	}

}
