package handler

import (
	"context"

	"github.com/Astemirdum/transactions/tx-user/internal/service"
	"github.com/Astemirdum/transactions/tx-user/internal/session"
	// _ "github.com/golang/mock/mockgen".
)

//go:generate go run github.com/golang/mock/mockgen -source=service.go -destination=mocks/mock.go

type UserService interface {
	SessionService
	CreateAccount(ctx context.Context, user *service.User) (*service.Account, error)
}

type SessionService interface {
	GenerateSessionID(ctx context.Context, login, password string) (*session.ID, error)
	CheckSessionID(ctx context.Context, id *session.ID) (*session.Session, error)
	DeleteSessionID(ctx context.Context, id *session.ID) error
}

var _ UserService = (*service.UserService)(nil)
