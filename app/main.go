package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/khusainnov/weather"
	"github.com/khusainnov/weather/pkg/handler"
	"github.com/khusainnov/weather/pkg/repository"
	"github.com/khusainnov/weather/pkg/service"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	logrus.Infoln("Reading config")
	if err := godotenv.Load(".env"); err != nil {
		logrus.Errorf("Cannot read .env file, due to error: %s", err.Error())
	}

	logrus.Infoln("Initializing repository")
	repos := repository.NewRepository()

	logrus.Infoln("Initializing services")
	services := service.NewService(repos)

	logrus.Infoln("Initializing handlers")
	handlers := handler.NewHandler(services)

	s := new(weather.Server)

	logrus.Infof("Starting server on port:%s", os.Getenv("PORT"))
	if err := s.Run(os.Getenv("PORT"), handlers.InitRoutesMux()); err != nil {
		logrus.Errorf("Cannot run server, due to error: %s", err.Error())
	}
}
