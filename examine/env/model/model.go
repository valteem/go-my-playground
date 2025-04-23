package model

import (
	"time"
)

type Config struct {
	App
	DB
	HTTP
}

type App struct {
	Version string `env:"APP_VERSION"`
}

type DB struct {
	MaxPoolSize int    `env:"DB_MAX_POOL_SIZE"`
	Url         string `env:"DB_URL"`
}

type HTTP struct {
	Port         string        `env:"HTTP_PORT"`
	ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT"`
	WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT"`
}
