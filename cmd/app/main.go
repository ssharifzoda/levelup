package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	handlers2 "github.com/ssharifzoda/levelup/internal/delivery/http/handlers"
	"github.com/ssharifzoda/levelup/internal/postgres"
	"github.com/ssharifzoda/levelup/internal/server"
	"github.com/ssharifzoda/levelup/internal/service"
	"github.com/ssharifzoda/levelup/pkg/database"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error initializing env value: %s", err.Error())
	}
	conn, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetUint16("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		logrus.Fatalf("failed to initializing db: %s", err.Error())
	}
	repository := database.NewDatabase(conn)
	services := service.NewService(repository)
	handlers := handlers2.NewHandler(services)
	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("cmd/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
