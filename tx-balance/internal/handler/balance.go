package handler

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	balancev1 "github.com/Astemirdum/transactions/proto/balance/v1"
	"github.com/Astemirdum/transactions/tx-balance/internal/handler/broker"
	models "github.com/Astemirdum/transactions/tx-balance/models/v1"
	"github.com/bufbuild/connect-go"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func (h *Handler) CreateBalance(
	ctx context.Context,
	req *connect.Request[balancev1.CreateBalanceRequest],
) (*connect.Response[balancev1.CreateBalanceResponse], error) {

	if err := h.svc.CreateBalance(ctx, int(req.Msg.UserId)); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	if err := h.mb.StartCashOutByUser(context.Background(), int(req.Msg.UserId), h.userCh); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&balancev1.CreateBalanceResponse{Cash: h.initCash}), nil
}

func (h *Handler) GetBalance(
	ctx context.Context,
	_ *connect.Request[balancev1.GetBalanceRequest],
) (*connect.Response[balancev1.GetBalanceResponse], error) {

	userID, err := getUserIDFromCtx(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}
	cash, err := h.svc.GetBalance(ctx, userID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	h.log.Debug("GetBalance", zap.Uint64("cash", cash))
	return connect.NewResponse(&balancev1.GetBalanceResponse{
		Cash: int64(cash),
	}), nil
}

func getUserIDFromCtx(ctx context.Context) (int, error) {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return 0, errors.New("no metadata")
	}
	v := md.Get(userIDKey)
	if len(v) == 0 {
		return 0, errors.New("no userID metadata")
	}
	return strconv.Atoi(v[0])
}

// CashOut is async method for cash out.
func (h *Handler) CashOut(
	ctx context.Context,
	req *connect.Request[balancev1.CashOutRequest],
) (*connect.Response[balancev1.CashOutResponse], error) {

	userID, err := getUserIDFromCtx(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}
	subject := fmt.Sprintf(broker.SubjectTmp, userID)
	if err := h.mb.PublishCashOut(subject, &models.CashOutMsg{
		Cash: uint64(req.Msg.Cash),
	}); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&balancev1.CashOutResponse{}), nil
}
