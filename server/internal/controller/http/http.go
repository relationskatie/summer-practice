package http

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/relationskatie/summer-practice/server/internal/config"
	"github.com/relationskatie/summer-practice/server/internal/controller"
	"github.com/relationskatie/summer-practice/server/internal/model"
	"github.com/relationskatie/summer-practice/server/internal/storage"
	"go.uber.org/zap"
)

var _ controller.Controller = (*Controller)(nil)

type Controller struct {
	server *echo.Echo
	log    *zap.Logger
	cfg    *config.Controller
	store  storage.Interface
	data   []model.FormResponse
	mutex  sync.Mutex
}

func NewServer(store storage.Interface, log *zap.Logger, cfg *config.Controller) (*Controller, error) {
	log.Info("Initialize controller")
	ctrl := &Controller{
		server: echo.New(),
		log:    log,
		cfg:    cfg,
		store:  store,
		data:   []model.FormResponse{},
		mutex:  sync.Mutex{},
	}
	ctrl.configure()
	return ctrl, nil
}

func (ctrl *Controller) configureRoutes() {
	log.Info("Configuration routes")
	api := ctrl.server.Group("/app")
	{
		api.GET("/", ctrl.HandleGetHomePage)
		api.GET("/form", ctrl.HandleGetForm)
		api.POST("/form", ctrl.HandlePostForm)
		api.GET("/vacancies", ctrl.HandleGetAllVacancies)
		api.GET("/vacancies/:id", ctrl.HandleGetVacancyByID)
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
	ctrl.log.Info("Start server", zap.String("bind-addres", ctrl.cfg.GetBindAddress()))
	return ctrl.server.Start(ctrl.cfg.GetBindAddress())
}

func (ctrl *Controller) ShutDown(ctx context.Context) error {
	ctrl.log.Info("Server ShutDown")
	return ctrl.server.Shutdown(ctx)
}
