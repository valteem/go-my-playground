package main

import (
	"github.com/caarlos0/env/v11"

	"github.com/valteem/examine/env/model"
)

func main() {

	LoadFromEnvFile(".env")

	cfg := model.Config{}

	env.Parse(&cfg)

}
