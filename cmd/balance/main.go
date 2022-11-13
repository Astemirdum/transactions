package main

import (
	"log"

	"github.com/Astemirdum/transactions/tx-balance/config"
	"github.com/Astemirdum/transactions/tx-balance/service"
	"github.com/spf13/pflag"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	configPath := pflag.StringP("config", "c", "balance.yml", "config path")
	pflag.Parse()

	if err := godotenv.Load(); err != nil {
		log.Fatal("load envs from .env", zap.Error(err))
	}
	cfg := config.GetConfigYML(*configPath)

	service.Run(cfg)
}
