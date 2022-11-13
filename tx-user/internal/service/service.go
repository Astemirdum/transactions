package service

import (
	"github.com/Astemirdum/transactions/tx-user/config"
	"github.com/Astemirdum/transactions/tx-user/internal/service/balanceclient"
	"go.uber.org/zap"

	"github.com/Astemirdum/transactions/tx-user/internal/repository"
	"github.com/Astemirdum/transactions/tx-user/internal/session"
	// _ "github.com/golang/mock/mockgen".
)

const (
	salt = "kjvbbe8392dsn"
)

type UserService struct {
	sm      *session.Manager
	balance *balanceclient.BalanceClient
	repo    UserRepo
	log     *zap.Logger
}

func NewUserService(
	repo *repository.UserDB,
	sm *session.Manager,
	cfg config.BalanceClient,
	log *zap.Logger,
) *UserService {

	return &UserService{
		repo:    repo,
		sm:      sm,
		balance: balanceclient.NewBalanceClient(cfg),
		log:     log,
	}
}
