package repository

import (
	"context"
	"database/sql"
	"errors"

	"go.uber.org/zap"

	models "github.com/Astemirdum/transactions/tx-balance/models/v1"
	"github.com/Masterminds/squirrel"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

var (
	ErrOverdraft = errors.New("not enough cash on bank account")
)

const balanceTable = "tv1.balances"

func (b *BalanceDB) CreateBalance(
	ctx context.Context,
	userID int,
) error {

	var id int
	query, args, err := psql.Insert(balanceTable).
		Values(userID, b.initCash).
		Columns("user_id", "cash").
		Suffix("RETURNING id").ToSql()
	if err != nil {
		return err
	}
	row := b.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&id); err != nil {
		return err
	}
	_ = id
	return err
}

func (b *BalanceDB) GetBalance(ctx context.Context, userID int) (uint64, error) {
	query, args, err := psql.Select("cash").
		From(balanceTable).
		Where("user_id=$1", userID).
		ToSql()
	if err != nil {
		return 0, err
	}
	var cash uint64
	row := b.db.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&cash); err != nil {
		return 0, err
	}
	return cash, nil
}

func (b *BalanceDB) CashOut(ctx context.Context, bl *models.Balance,
) (int64, error) {

	tx, err := b.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			if txErr := tx.Rollback(); txErr != nil {
				b.log.Debug("tx failed", zap.Error(txErr))
			}
		}
	}()

	query, args, err := psql.Select("cash").
		From(balanceTable).
		Where("user_id=$1", bl.UserID).
		ToSql()
	if err != nil {
		b.log.Error("cash out", zap.String("query", query), zap.Any("args", args))
		return 0, err
	}
	row := tx.QueryRowContext(ctx, query, args...)
	var cash int64
	if err := row.Scan(&cash); err != nil {
		return 0, err
	}
	remain := cash - int64(bl.Cash)
	if remain < 0 {
		remain = 0
	}
	query, args, err = psql.
		Update(balanceTable).
		Set("cash", remain).
		Where("user_id=$2", bl.UserID).
		ToSql()
	if err != nil {
		b.log.Error("cash out", zap.String("query", query), zap.Any("args", args))
		return 0, err
	}
	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		return 0, err
	}

	if cash < int64(bl.Cash) {
		_ = tx.Commit() //nolint:errcheck
		return 0, ErrOverdraft
	}

	return remain, tx.Commit()
}

func (b *BalanceDB) ListCashedBalance(ctx context.Context) ([]models.Balance, error) {
	query, args, err := psql.Select("user_id", "cash").
		From(balanceTable).
		Where("cash > 0").
		ToSql()
	if err != nil {
		b.log.Error("cash out", zap.String("query", query), zap.Any("args", args))
		return nil, err
	}
	var balances []models.Balance
	if err := b.db.SelectContext(ctx, &balances, query, args...); err != nil {
		return nil, err
	}
	return balances, nil
}
