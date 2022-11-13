package service

import (
	"context"
	"database/sql"

	"github.com/Astemirdum/transactions/tx-user/internal/repository"
	models "github.com/Astemirdum/transactions/tx-user/models/v1"
	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user models.User, tx *sqlx.Tx) (int, error)
	GetUser(ctx context.Context, email, password string) (models.User, error)
	ExistsEmail(ctx context.Context, email string) (bool, error)
	InTx(ctx context.Context, opts *sql.TxOptions, fn func(tx *sqlx.Tx) error) error
}

var _ UserRepo = (*repository.UserDB)(nil)
