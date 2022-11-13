package handler

import (
	"context"
	"errors"
	"strconv"

	"github.com/Astemirdum/transactions/pkg"
	"github.com/Astemirdum/transactions/proto/balance/v1/balancev1connect"
	"github.com/bufbuild/connect-go"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

const (
	userIDKey              = "userid"
	createBalanceProcedure = "/" + balancev1connect.BalanceServiceName + "/CreateBalance"
)

func (h *Handler) AuthInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {

			if req.Spec().Procedure == createBalanceProcedure {
				return next(ctx, req)
			}

			session := req.Header().Get(pkg.SessionKey)
			if session == "" {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("no session provided"),
				)
			}
			userID, err := h.auth.Auth(ctx, session)
			if err != nil {
				return nil, connect.NewError(connect.CodeOf(err), err)
			}

			ctx = metadata.AppendToOutgoingContext(ctx, userIDKey, strconv.Itoa(int(userID)))
			h.log.Debug("AuthInterceptor", zap.Int32("userID", userID))
			return next(ctx, req)
		}
	}
}

func (h *Handler) LoggingInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {

			h.log.Info("LoggingInterceptor",
				zap.String("procedure", req.Spec().Procedure),
				zap.Bool("isClient", req.Spec().IsClient),
			)
			return next(ctx, req)
		}
	}
}

func (h *Handler) ValidateInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {

			if err := validateRequest(req.Any().(Validate)); err != nil {
				return nil, connect.NewError(connect.CodeInvalidArgument, err)
			}
			return next(ctx, req)
		}
	}
}

type Validate interface {
	Validate() error
}

func validateRequest[T Validate](msg T) error {
	return msg.Validate()
}
