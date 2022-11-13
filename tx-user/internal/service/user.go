package service

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"github.com/Astemirdum/transactions/tx-user/internal/session"
	"github.com/jmoiron/sqlx"

	models "github.com/Astemirdum/transactions/tx-user/models/v1"
)

var (
	ErrNoUser             = errors.New("no user exists")
	ErrEmailAlreadyExists = errors.New("email already exists")
)

type Account struct {
	UserID int
	Cash   int64
}

type User struct {
	Email    string
	Password string
}

func (a *UserService) CreateAccount(ctx context.Context, user *User) (*Account, error) {
	exists, err := a.repo.ExistsEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrEmailAlreadyExists
	}

	var acc Account
	if err := a.repo.InTx(ctx, nil, func(tx *sqlx.Tx) error {
		if acc.UserID, err = a.repo.CreateUser(ctx, models.User{
			Password: genPasswordHash(user.Password),
			Email:    user.Email,
		}, tx); err != nil {
			a.log.Error("CreateBalance", zap.Error(err))
			return err
		}
		if acc.Cash, err = a.balance.CreateBalance(ctx, acc.UserID); err != nil {
			a.log.Error("CreateBalance", zap.Error(err))
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &acc, nil
}

func (a *UserService) DeleteSessionID(ctx context.Context, id *session.ID) error {
	return a.sm.Delete(ctx, id)
}

func (a *UserService) CheckSessionID(ctx context.Context, id *session.ID) (*session.Session, error) {
	return a.sm.Check(ctx, id)
}

func (a *UserService) GenerateSessionID(ctx context.Context, login, password string) (*session.ID, error) {
	user, err := a.repo.GetUser(ctx, login, genPasswordHash(password))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoUser
		}
		return nil, err
	}
	return a.sm.Create(ctx, &session.Session{
		Login:  login,
		UserID: user.ID,
	})
}

func genPasswordHash(pass string) string {
	hash := sha256.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
