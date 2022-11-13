package handler

import (
	"context"

	"github.com/bufbuild/connect-go"
	"go.uber.org/zap"
)

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
