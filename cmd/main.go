package main

import (
	grpc_client "github.com/KonstantinP85/grpc-client"
	"github.com/KonstantinP85/grpc-client/pkg/handler"
	"github.com/KonstantinP85/grpc-client/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init config: %s", err.Error())
	}

	services := service.NewService()
	handlers := handler.NewHandler(services)

	srv := new(grpc_client.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("server run error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
