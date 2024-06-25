package http

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	if err := ctrl.configure(); err != nil {
		return nil, err
	}
	return ctrl, nil
}

func (ctrl *Controller) configureRoutes() {
	ctrl.server.GET("/hi", ctrl.HadleHi)
}

func (ctrl *Controller) configureMiddlewares() {
	var middlewares = []echo.MiddlewareFunc{
		middleware.Gzip(),
		middleware.Recover(),
		middleware.RequestIDWithConfig(middleware.RequestIDConfig{
			Skipper:      middleware.DefaultSkipper,
			Generator:    uuid.NewString,
			TargetHeader: echo.HeaderXRequestID,
		}),
		middleware.Logger(),
	}
	ctrl.server.Use(middlewares...)
}

func (ctrl *Controller) configure() error {
	ctrl.configureMiddlewares()
	ctrl.configureRoutes()
	return nil
}

func (ctrl *Controller) Start(ctx context.Context) error {
	return ctrl.server.Start(":8080")
}

func (ctrl *Controller) ShutDown(ctx context.Context) error {
	return ctrl.server.Shutdown(ctx)
}
