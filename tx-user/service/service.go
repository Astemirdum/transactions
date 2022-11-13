package service

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Astemirdum/transactions/tx-user/config"
	"github.com/Astemirdum/transactions/tx-user/internal/handler"
	"github.com/Astemirdum/transactions/tx-user/internal/repository"
	"github.com/Astemirdum/transactions/tx-user/internal/service"
	"github.com/Astemirdum/transactions/tx-user/internal/session"
	"go.uber.org/zap"
)

func Run(cfg *config.Config) {
	log := zap.NewExample()
	db, err := repository.NewDB(&cfg.Database)
	if err != nil {
		log.Fatal("postgres db conn", zap.Error(err))
	}

	ctx := context.Background()
	sm, err := session.NewManager(ctx, cfg.Redis)
	if err != nil {
		log.Fatal("redis init", zap.Error(err))
	}

	// repo -> service -> handler
	repo := repository.NewUserDB(db)
	services := service.NewUserService(repo, sm, cfg.BalanceClient, log.Named("svc"))
	handlers := handler.NewHandler(services, log.Named("handler"))

	addr := net.JoinHostPort(cfg.Server.Host, cfg.Server.Port)
	srv := handler.NewServer(handlers.NewRouter(),
		handler.WithAddrOption(addr),
		handler.WithTimeOutOption(cfg.Server.Timeout))
	go func() {
		if err := srv.Run(); err != nil {
			log.Fatal("Server init", zap.Error(err))
		}
	}()
	log.Info("server has been started on", zap.String("addr", addr))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Info("Graceful shutdown")
	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()
	if err = srv.Shutdown(ctx); err != nil {
		log.Error("server Shutdown", zap.Error(err))
	}
	if err = db.Close(); err != nil {
		log.Error("db connection close", zap.Error(err))
	}
}
