package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/ssharifzoda/levelup/internal/database"
	handlers2 "github.com/ssharifzoda/levelup/internal/delivery/http/handlers"
	"github.com/ssharifzoda/levelup/internal/server"
	"github.com/ssharifzoda/levelup/internal/service"
	"github.com/ssharifzoda/levelup/pkg/logging"
	"os"
)

func main() {
	logger := logging.GetLogger()
	if err := initConfig(); err != nil {
		logger.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error initializing env value: %s", err.Error())
	}
	conn, err := database.NewPostgresDB(database.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetUint16("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		logger.Fatalf("failed to initializing db: %s", err.Error())
	}
	session, err := database.NewPostgresGorm()
	if err != nil {
		logger.Fatalf("failed to initializing db: %s", err.Error())
	}
	repository := database.NewDatabase(conn, session)
	services := service.NewService(repository)
	handlers := handlers2.NewHandler(services)
	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logger.Fatalf("error occured while running http server: %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("internal/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
