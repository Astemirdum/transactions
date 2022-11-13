package handler

import (
	"context"

	"github.com/Astemirdum/transactions/tx-balance/internal/handler/broker"
	"github.com/Astemirdum/transactions/tx-balance/internal/service"

	models "github.com/Astemirdum/transactions/tx-balance/models/v1"
	// _ "github.com/golang/mock/mockgen".
)

//go:generate go run github.com/golang/mock/mockgen -source=service.go -destination=mocks/mock.go

type Balance interface {
	CreateBalance(ctx context.Context, userID int) error
	CashOut(ctx context.Context, msg *models.CashOut) (int64, error)
	GetBalance(ctx context.Context, userID int) (uint64, error)
	ListCashedUserIDs(ctx context.Context) ([]int, error)
}

var _ Balance = (*service.BalanceService)(nil)

type Auth interface {
	Auth(ctx context.Context, session string) (int32, error)
}

type CashOutBroker interface {
	SetCashOutHandler(cashOut broker.CashOut)
	StartCashOutByUser(ctx context.Context, userID int, unSubCh chan int) error
	UnsubscribeUser(userID int) error
	PublishCashOut(subjectName string, cashMsg *models.CashOutMsg) error
}
