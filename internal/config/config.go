package config

import (
	"context"
	"os"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/flags"
)

type Config struct {
	Controller *Controller
	Postgres   *PostgresConfig
}

func New(ctx context.Context) (*Config, error) {
	_, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	cfg := &Config{
		Controller: &Controller{
			BindAddres: "localhost",
			BindPort:   8080,
		},
		Postgres: &PostgresConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "postgres",
			Database: "postgres",
		},
	}
	loader := confita.NewLoader(env.NewBackend(), flags.NewBackend())
	if err = loader.Load(ctx, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
