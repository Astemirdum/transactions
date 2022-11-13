package repository

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"

	"github.com/jmoiron/sqlx"
)

type UserDB struct {
	db *sqlx.DB
}

func NewUserDB(db *sqlx.DB) *UserDB {
	return &UserDB{db: db}
}

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

const (
	userTable = "tv1.users"
)

func (u *UserDB) InTx(
	ctx context.Context,
	opts *sql.TxOptions,
	fn func(tx *sqlx.Tx) error,
) error {

	tx, err := u.db.BeginTxx(ctx, opts)
	if err != nil {
		return err
	}
	if err := fn(tx); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			zap.L().Error("tx failed", zap.Error(txErr))
		}
		return err
	}

	return tx.Commit()
}
