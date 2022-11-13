package service

import (
	"context"

	"github.com/Astemirdum/transactions/pkg"
	models "github.com/Astemirdum/transactions/tx-balance/models/v1"
)

func (a *BalanceService) CreateBalance(ctx context.Context, userID int) error {
	return a.repo.CreateBalance(ctx, userID)
}

func (a *BalanceService) GetBalance(ctx context.Context, userID int) (uint64, error) {
	return a.repo.GetBalance(ctx, userID)
}

func (a *BalanceService) ListCashedUserIDs(ctx context.Context) ([]int, error) {
	list, err := a.repo.ListCashedBalance(ctx)
	if err != nil {
		return nil, err
	}
	return pkg.Map(list, func(a models.Balance) int { return a.UserID }), nil
}

func (a *BalanceService) CashOut(ctx context.Context, msg *models.CashOut,
) (int64, error) {

	remainCash, err := a.repo.CashOut(ctx, &models.Balance{
		UserID: msg.UserID,
		Cash:   msg.Cash,
	})
	if err != nil {
		return 0, err
	}
	return remainCash, nil
}
