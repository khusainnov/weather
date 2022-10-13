package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/khusainnov/weather"
	"github.com/khusainnov/weather/pkg/handler"
	"github.com/khusainnov/weather/pkg/repository"
	"github.com/khusainnov/weather/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	logrus.Infoln("Reading config from .env")
	if err := godotenv.Load("./config/.env"); err != nil {
		logrus.Errorf("Cannot read .env file, due to error: %s", err.Error())
	}

	logrus.Infoln("Reading config from .yml")
	if err := initConfig(); err != nil {
		logrus.Errorf("Cannot reading config from .yml, due to error: %s", err.Error())
	}

	logrus.Infoln("Initializing DB")
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USER"),
		DBName:   os.Getenv("PG_DB_NAME"),
		SSLMode:  os.Getenv("PG_SSL_MODE"),
		Password: os.Getenv("PG_PASSWORD"),
	})

	/*db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.user"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})*/

	if err != nil {
		logrus.Errorf("Cannot run db, due to error: %s", err.Error())
	}

	logrus.Infoln("Initializing repository")
	repos := repository.NewRepository(db)

	logrus.Infoln("Initializing services")
	services := service.NewService(repos)

	logrus.Infoln("Initializing handlers")
	handlers := handler.NewHandler(services)

	s := new(weather.Server)

	logrus.Infof("Starting server on port:%s", os.Getenv("PORT"))
	if err = s.Run(os.Getenv("PORT"), handlers.InitRoutesMux()); err != nil {
		logrus.Errorf("Cannot run server, due to error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
