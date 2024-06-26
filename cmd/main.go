package main

import (
	"context"

	"github.com/relationskatie/summer-practice/internal/config"
	"github.com/relationskatie/summer-practice/internal/controller"
	"github.com/relationskatie/summer-practice/internal/controller/http"
	"go.uber.org/zap"
)

func main() {
	var (
		log    *zap.Logger
		err    error
		ctx    context.Context
		server controller.Controller
		cnf    *config.Config
	)

	log, err = zap.NewProduction()
	if err != nil {
		log.Fatal("Failed to initialize logger", zap.Error(err))
	}

	log.Info("Initialized loger")

	server, err = http.NewServer(log, cnf)
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
