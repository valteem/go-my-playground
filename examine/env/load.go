package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadFromEnvFile(path string, envs ...string) {

	dir, _ := os.Getwd()
	log.Printf("%s", dir)

	err := godotenv.Load(path)
	if err != nil {
		log.Fatalf("failed to load env variables from %s path: %v", path, err)
	}

	for _, env := range envs {
		log.Printf("%s = %s", env, os.Getenv(env))
	}

}
