package http

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
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
	log.Info("Configuration routes")
	api := ctrl.server.Group("/app")
	{
		vacancy := api.Group("/vacans")
		{
			vacancy.POST("/", ctrl.HandleGetVacanciesByTunning)
			vacancy.GET("/:id", ctrl.HandleGetVacancyById)
			vacancy.POST("/:id", ctrl.HandleAddToFavourite)
			vacancy.GET("/fav/", ctrl.HandleGetAllFavourite)
			vacancy.DELETE("/fav/:id", ctrl.HandleDeleteIntoFavourite)

		}
	}
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
	ctrl.log.Info("Start server on port :8080")
	return ctrl.server.Start(":8080")
}

func (ctrl *Controller) ShutDown(ctx context.Context) error {
	ctrl.log.Info("Server ShutDown")
	return ctrl.server.Shutdown(ctx)
}
