package config

import (
	"context"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/flags"
)

type Config struct {
	Controller *Controller
	//Postgres   *PostgresConfig
}

func New(ctx context.Context) (*Config, error) {
	cfg := &Config{
		Controller: new(Controller),
	}
	loader := confita.NewLoader(env.NewBackend(), flags.NewBackend())
	if err := loader.Load(ctx, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
