package main

import (
	"github.com/spf13/viper"
	"log"
	"rest/pkg/handler"
	"rest/pkg/httpserver"
	"rest/pkg/repository"
	"rest/pkg/service"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("pizda configu: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(httpserver.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("pizdec: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
