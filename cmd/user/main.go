package main

import (
	"github.com/Astemirdum/transactions/tx-user/config"
	"github.com/Astemirdum/transactions/tx-user/service"
	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
)

func main() {
	configPath := pflag.StringP("config", "c", "user.yml", "config path")
	pflag.Parse()
	log := zap.NewExample()

	if err := godotenv.Load(); err != nil {
		log.Fatal("load envs from .env", zap.Error(err))
	}

	cfg := config.GetConfigYML(*configPath)
	service.Run(cfg)
}
