package app

import (
	"github.com/labstack/echo/v4"
	"github.com/rostis232/parcelstrackingservice/internal/handler"
	"github.com/rostis232/parcelstrackingservice/internal/routes"
	"github.com/rostis232/parcelstrackingservice/internal/service"
	"github.com/rostis232/parcelstrackingservice/models"
	"github.com/sirupsen/logrus"
)

type App struct {
	config *models.AppConfig
	Server *echo.Echo
}

func New(config *models.AppConfig) *App {
	a := &App{
		config: config,
	}
	a.Server = echo.New()
	routes.RegisterRoutes(a.Server, handler.New(service.New(config.QueueTimeOut, config.MaxQueueCount)))
	return a
}

func (a *App) Start() {
	logrus.Fatal(a.Server.Start(":" + a.config.Port))
}
