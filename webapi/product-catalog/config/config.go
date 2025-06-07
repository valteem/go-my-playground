package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	HTTP
	PG
	JWT
}

type HTTP struct {
	Port string `env:"PRODUCT_CATALOG_HTTP_PORT"`
}

type PG struct {
	MaxPoolSize int    `env:"PRODUCT_CATALOG_PG_MAX_POOL_SIZE"`
	URL         string `env:"PRODUCT_CATALOG_PG_URL"`
}

type JWT struct {
	SignKey string `env:"PRODUCT_CATALOG_JWT_SIGN_KEY"`
}

func Load(ctx context.Context) (*Config, error) {

	cfg := &Config{}

	err := envconfig.Process(ctx, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil

}
