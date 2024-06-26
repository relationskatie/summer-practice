package main

import (
	"context"

	"github.com/relationskatie/summer-practice/server/internal/config"
	"github.com/relationskatie/summer-practice/server/internal/controller"
	"github.com/relationskatie/summer-practice/server/internal/controller/http"
	"go.uber.org/zap"
)

func main() {
	var (
		log    *zap.Logger
		err    error
		ctx    context.Context
		server controller.Controller
		cfg    *config.Config
	)
	ctx = context.Background()
	log, err = zap.NewProduction()
	if err != nil {
		log.Fatal("Failed to initialize logger", zap.Error(err))
	}

	log.Info("Initialized loger")

	cfg, err = config.New(ctx)
	if err != nil {
		log.Fatal("Failed to initialize configuration", zap.Error(err))
	}
	log.Info("initialized configuration", zap.Any("cnf", cfg))

	server, err = http.NewServer(log)
	if err != nil {
		log.Fatal("Failed to initialize server", zap.Error(err))
	}
	defer func() {
		log.Error(
			"stopped server",
			zap.Error(server.ShutDown(ctx)),
		)
	}()
	err = server.Start(ctx)
	if err != nil {
		log.Error("Failed to start server")
	}
}
