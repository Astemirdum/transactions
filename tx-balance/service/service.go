package service

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Astemirdum/transactions/tx-balance/config"
	"github.com/Astemirdum/transactions/tx-balance/internal/handler"
	"github.com/Astemirdum/transactions/tx-balance/internal/handler/broker"
	"github.com/Astemirdum/transactions/tx-balance/internal/repository"
	"github.com/Astemirdum/transactions/tx-balance/internal/service"
	"go.uber.org/zap"
)

func Run(cfg *config.Config) {
	log := zap.NewExample()
	db, err := repository.NewDB(&cfg.Database)
	if err != nil {
		log.Fatal("postgres db conn", zap.Error(err))
	}

	br, err := broker.NewBroker(cfg.JS, log.Named("js"))
	if err != nil {
		log.Fatal("broker init", zap.Error(err))
	}
	// repo -> service -> handler
	repo := repository.NewBalanceDB(db, log.Named("db"), cfg.InitCash)
	services := service.NewBalanceService(repo)
	handlers := handler.NewHandler(services, br, log.Named("handler"), cfg)
	addr := net.JoinHostPort(cfg.Server.Host, cfg.Server.Port)
	srv := handler.NewServer(handlers.NewRouter(), addr, cfg.Server)
	go func() {
		if err := srv.Run(); err != nil {
			log.Fatal("Server init", zap.Error(err))
		}
	}()
	log.Info("Server has been started on", zap.String("addr", addr))

	if err := handlers.StartCashing(context.Background()); err != nil {
		log.Fatal("startCashing", zap.Error(err))
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Info("Graceful shutdown")
	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second)
	defer cancelFn()
	if err = srv.Shutdown(ctx); err != nil {
		log.Error("server Shutdown", zap.Error(err))
	}
	if err = db.Close(); err != nil {
		log.Error("db connection close", zap.Error(err))
	}
	br.Close()
}
