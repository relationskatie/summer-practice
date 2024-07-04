package config

import (
	"context"
	"log"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/flags"
)

type Config struct {
	Controller *Controller
	Postgres   *PostgresConfig
}

func New(ctx context.Context) (*Config, error) {
	cfg := &Config{
		Controller: &Controller{
			BindAddres: "localhost",
			BindPort:   8000,
		},
		Postgres: &PostgresConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "postgres",
			Database: "vacancies",
		},
	}
	loader := confita.NewLoader(env.NewBackend(), flags.NewBackend())
	if err := loader.Load(ctx, cfg); err != nil {
		log.Fatal("Failed to loader load")
		return nil, err
	}
	return cfg, nil
}
