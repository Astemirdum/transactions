package handler

import (
	"context"
	"net/http"

	"github.com/Astemirdum/transactions/proto/balance/v1/balancev1connect"
	"github.com/Astemirdum/transactions/tx-balance/config"
	"github.com/Astemirdum/transactions/tx-balance/internal/handler/broker"
	"github.com/Astemirdum/transactions/tx-balance/internal/handler/userclient"
	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Handler struct {
	svc      Balance
	mb       CashOutBroker // *broker.Broker
	auth     Auth          // *userclient.AuthClient
	log      *zap.Logger
	userCh   chan int
	initCash int64
}

func NewHandler(
	srv Balance,
	br *broker.Broker,
	logger *zap.Logger,
	cfg *config.Config,
) *Handler {

	h := &Handler{
		svc:      srv,
		mb:       br,
		log:      logger,
		auth:     userclient.NewAuthClient(cfg.AuthClient),
		userCh:   make(chan int),
		initCash: int64(cfg.InitCash),
	}
	return h
}

func (h *Handler) StartCashing(ctx context.Context) error {
	h.mb.SetCashOutHandler(h.svc.CashOut)
	ids, err := h.svc.ListCashedUserIDs(ctx)
	if err != nil {
		return err
	}
	for _, id := range ids {
		if err := h.mb.StartCashOutByUser(ctx, id, h.userCh); err != nil {
			return err
		}
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case userID := <-h.userCh:
				if err := h.mb.UnsubscribeUser(userID); err != nil {
					h.log.Debug("unsubscribeUser",
						zap.Int("userID", userID))
				}
			}
		}
	}()
	return nil
}

func (h *Handler) NewRouter() http.Handler {
	mux := http.NewServeMux()
	{
		api := http.NewServeMux()
		api.Handle(balancev1connect.NewBalanceServiceHandler(h,
			connect.WithInterceptors(h.LoggingInterceptor()),
			connect.WithInterceptors(h.ValidateInterceptor()),
			connect.WithInterceptors(h.AuthInterceptor()),
		))
		mux.Handle("/grpc/balance/", http.StripPrefix("/grpc/balance", api))

		compress1KB := connect.WithCompressMinBytes(1024)
		mux.Handle(grpchealth.NewHandler(
			grpchealth.NewStaticChecker(balancev1connect.BalanceServiceName),
			compress1KB,
		))
		mux.Handle(grpcreflect.NewHandlerV1(
			grpcreflect.NewStaticReflector(balancev1connect.BalanceServiceName),
			compress1KB,
		))
		mux.Handle(grpcreflect.NewHandlerV1Alpha(
			grpcreflect.NewStaticReflector(balancev1connect.BalanceServiceName),
			compress1KB,
		))
	}

	return h2c.NewHandler(
		mux,
		&http2.Server{})
}

func NewHandlerTest(
	srv Balance,
	br CashOutBroker,
	auth Auth,
	logger *zap.Logger,
	cfg *config.Config,
) *Handler {

	h := &Handler{
		svc:      srv,
		mb:       br,
		log:      logger,
		auth:     auth,
		userCh:   make(chan int),
		initCash: int64(cfg.InitCash),
	}
	return h
}
