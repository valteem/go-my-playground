package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	HTTP `yaml:"http"`
	PG   `yaml:"postgres"`
	JWT  `yaml:"jwt"`
}

type HTTP struct {
	Port string `yaml:"http_port"`
}

type PG struct {
	MaxPoolSize int    `yaml:"max_pool_size"`
	URL         string `yaml:"url"`
}

type JWT struct {
	SignKey string `yaml:"jwt_sign_key"`
}

func Load(s string) (*Config, error) {

	cfg := &Config{}

	err := yaml.Unmarshal([]byte(s), cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil

}

func LoadFile(path string) (*Config, error) {

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg, err := Load(string(b))

	return cfg, err

}
