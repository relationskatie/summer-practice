package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/relationskatie/summer-practice/server/internal/config"
	"github.com/relationskatie/summer-practice/server/internal/controller/http"
	"github.com/relationskatie/summer-practice/server/internal/storage/pgx"
	"github.com/relationskatie/summer-practice/server/pgx/postgres"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	log, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Failed to initialize logger", zap.Error(err))
	}
	defer log.Sync()
	log.Info("Initialized loger")

	cfg, err := config.New(ctx)
	if err != nil {
		log.Fatal("Failed to initialize configuration", zap.Error(err))
	}
	log.Info("initialized configuration", zap.Any("cnf", cfg))

	pool, err := postgres.New(cfg, log)
	if err != nil {
		log.Fatal("Failed to initialize pool", zap.Error(err))
	}
	var pgErr *pgconn.PgError
	store, err := pgx.New(pool, log, pgErr)
	if err != nil {
		log.Fatal("Failed to create pgx storage", zap.Error(err))
	}
	server, err := http.NewServer(log, cfg.Controller)
	if err != nil {
		log.Fatal("Failed to initialize server", zap.Error(err))
	}
	defer func() {
		if err := server.ShutDown(ctx); err != nil {
			log.Error("Failed to shut down server", zap.Error(err))
		} else {
			log.Info("Server shut down gracefully")
		}
	}()
	err = server.Start(ctx)
	if err != nil {
		log.Error("Failed to start server")
	}
}
