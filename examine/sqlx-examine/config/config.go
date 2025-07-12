package config

import (
	"fmt"
	"os"
)

type Config struct {
	PgURL string
}

func NewConfig() (*Config, error) {

	pgURL := os.Getenv("PG_URL")
	if pgURL == "" {
		return nil, fmt.Errorf("empty database URL")
	}

	return &Config{PgURL: pgURL}, nil

}
