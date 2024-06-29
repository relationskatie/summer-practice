package main

import (
	"context"

	"github.com/relationskatie/summer-practice/server/internal/config"
	"github.com/relationskatie/summer-practice/server/internal/controller/http"
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
