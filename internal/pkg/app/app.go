package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

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
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
        if err := a.Server.Start(":" + a.config.Port); err != nil && err != http.ErrServerClosed {
            logrus.Fatalf("shutting down the server: %v", err)
        }
    }()

	<-ctx.Done()
	logrus.Println("Shutting down gracefully, press Ctrl+C again to force")
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {
        logrus.Fatalf("server forced to shutdown: %v", err)
    }
}
