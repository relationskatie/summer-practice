package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/relationskatie/summer-practice/server/internal/config"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
	"go.uber.org/zap"
)

func New(cfg *config.Config, log *zap.Logger) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool
	log.Info("Initializing postgres client")
	p, err := pgxpool.ParseConfig(cfg.Postgres.DataBaseDNS())
	if err != nil {
		return nil, fmt.Errorf("Failed to parse db: %w", err)
	}
	p.AfterConnect = func(ctx context.Context, c *pgx.Conn) error {
		pgxUUID.Register(c.TypeMap())
		return nil
	}
	log.Info("Created postgres client")
	pool, err = pgxpool.NewWithConfig(context.Background(), p)
	if err != nil {
		return nil, fmt.Errorf("postgres: init pgxpool: %w", err)
	}
	log.Info("created postgres client")
	return pool, nil
}
