package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ServerPort string `env:"SERVER_PORT"`

	User    string `env:"PG_USER"`
	Pass    string `env:"PG_PASS"`
	Host    string `env:"PG_HOST"`
	Port    string `env:"PG_PORT"`
	Name    string `env:"PG_DB"`
	SSLMode string `env:"PG_SSLMODE"`
}

func Get() (*Config, error) {
	var cfg Config

	if err := godotenv.Load(); err == nil {
		logrus.Info(".env loaded")
	}

	err := env.Parse(&cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse config")
	}
	return &cfg, nil
}
