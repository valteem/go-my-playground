package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type JWT struct {
	Secret string `env:"JWT_SECRET"`
	Expire int    `env:"JWT_EXPIRE"` // token expiration time in hours
}

type AdminUser struct {
	UserName string `env:"ADM_NAME"`
	Password string `env:"ADM_PASSWD"`
}

type Config struct {
	JWT
	AdminUser
}

func Load(ctx context.Context) (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process(ctx, cfg)
	return cfg, err
}
