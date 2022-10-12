package main

import (
	"github.com/spf13/viper"
	"github.com/ssharifzoda/levelup/internal/delivery/http/handler"
	"github.com/ssharifzoda/levelup/internal/server"
	"github.com/ssharifzoda/levelup/internal/service"
	"github.com/ssharifzoda/levelup/pkg/database"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	db := database.NewDatabase()
	services := service.NewService(db)
	handlers := handler.NewHandler(services)
	srv := new(server.Server)
	if err := srv.Run(viper.GetString("9999"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("cmd/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
