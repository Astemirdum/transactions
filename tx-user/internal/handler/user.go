package handler

import (
	"context"
	"errors"

	"github.com/Astemirdum/transactions/pkg"
	userv1 "github.com/Astemirdum/transactions/proto/user/v1"
	"github.com/Astemirdum/transactions/tx-user/internal/service"
	"github.com/Astemirdum/transactions/tx-user/internal/session"
	"github.com/bufbuild/connect-go"
	"go.uber.org/zap"
)

func (h *Handler) SignUp(ctx context.Context, req *connect.Request[userv1.SignUpRequest],
) (*connect.Response[userv1.SignUpResponse], error) {

	acc, err := h.svc.CreateAccount(ctx, &service.User{
		Email:    req.Msg.User.Email,
		Password: req.Msg.User.Password,
	})
	if err != nil {
		h.log.Error("createAccount", zap.Error(err))
		if errors.Is(err, service.ErrEmailAlreadyExists) {
			return nil, connect.NewError(connect.CodeAlreadyExists, err)
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res := connect.NewResponse(&userv1.SignUpResponse{
		UserId:  int32(acc.UserID),
		Balance: &userv1.Balance{Cash: acc.Cash},
	})
	res.Header().Set("User-Version", "v1")
	return res, nil
}

func (h *Handler) Auth(ctx context.Context, req *connect.Request[userv1.AuthRequest],
) (*connect.Response[userv1.AuthResponse], error) {

	ss, err := h.svc.CheckSessionID(ctx, &session.ID{
		ID: req.Msg.SessionId.Id,
	})
	if err != nil {
		if errors.Is(err, session.ErrSessionNotFound) {
			return nil, connect.NewError(connect.CodeUnauthenticated, err)
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&userv1.AuthResponse{
		UserId: int32(ss.UserID),
	}), nil
}

func (h *Handler) SignIn(ctx context.Context, req *connect.Request[userv1.SignInRequest],
) (*connect.Response[userv1.SignInResponse], error) {

	user := req.Msg.User
	id, err := h.svc.GenerateSessionID(ctx, user.Email, user.Password)
	if err != nil {
		if errors.Is(err, service.ErrNoUser) {
			return nil, connect.NewError(connect.CodeUnauthenticated, err)
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	h.log.Debug("signIn user",
		zap.String("email", user.Email),
		zap.String("sessionID", id.ID))

	resp := connect.NewResponse(&userv1.SignInResponse{
		SessionId: &userv1.SessionID{
			Id: id.ID,
		},
	})
	resp.Header().Set(pkg.SessionKey, id.ID)
	return resp, nil
}

func (h *Handler) SignOut(ctx context.Context, req *connect.Request[userv1.SignOutRequest],
) (*connect.Response[userv1.SignOutResponse], error) {

	if err := h.svc.DeleteSessionID(ctx, &session.ID{ID: req.Msg.SessionId.Id}); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&userv1.SignOutResponse{}), nil
}
