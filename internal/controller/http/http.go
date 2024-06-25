package http

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/relationskatie/summer-practice/internal/controller"
	"go.uber.org/zap"
)

var _ controller.Controller = (*Controller)(nil)

type Controller struct {
	server *echo.Echo
	log    *zap.Logger
}

func NewServer(log *zap.Logger) (*Controller, error) {
	log.Info("Initialize controller")
	ctrl := &Controller{
		server: echo.New(),
		log:    log,
	}
	return ctrl, nil
}

func (ctrl *Controller) Start(ctx context.Context) error {
	return ctrl.server.Start(":8080")
}

func (ctrl *Controller) ShutDown(ctx context.Context) error {
	return ctrl.server.Shutdown(ctx)
}
