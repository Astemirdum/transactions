package service

import (
	"context"

	"github.com/Astemirdum/transactions/tx-balance/internal/repository"
	models "github.com/Astemirdum/transactions/tx-balance/models/v1"
)

type Balance interface {
	CashOut(ctx context.Context, balance *models.Balance) (int64, error)
	CreateBalance(ctx context.Context, userID int) error
	GetBalance(ctx context.Context, userID int) (uint64, error)
	ListCashedBalance(ctx context.Context) ([]models.Balance, error)
}

var _ Balance = (*repository.BalanceDB)(nil)
