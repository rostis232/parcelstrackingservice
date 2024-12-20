package main

import (
	"github.com/rostis232/parcelstrackingservice/internal/pkg/app"
	"github.com/rostis232/parcelstrackingservice/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	err := initConfig()
	if err != nil {
		logrus.Fatalf("Error occured while initializing config: %s", err)
	}

	timeOutText := viper.GetString("timeout")
	timeOut, err := time.ParseDuration(timeOutText)
	if err != nil {
		logrus.Fatalf("Error occured while parsing timeout from config: %s", err)
	}
	config := models.AppConfig{
		Port:          viper.GetString("port"),
		MaxQueueCount: viper.GetInt("max-queue-count"),
		QueueTimeOut:  timeOut,
	}

	a := app.New(&config)
	a.Start()
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
