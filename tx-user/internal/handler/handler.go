package handler

import (
	"net/http"

	"github.com/Astemirdum/transactions/proto/user/v1/userv1connect"
	"github.com/Astemirdum/transactions/tx-user/internal/service"
	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Handler struct {
	svc UserService
	log *zap.Logger
}

func NewHandler(
	srv *service.UserService,
	logger *zap.Logger,
) *Handler {

	return &Handler{
		svc: srv,
		log: logger,
	}
}

func (h *Handler) NewRouter() http.Handler {
	mux := http.NewServeMux()
	compress1KB := connect.WithCompressMinBytes(1024)

	api := http.NewServeMux()
	api.Handle(userv1connect.NewUserServiceHandler(h,
		connect.WithInterceptors(h.LoggingInterceptor()),
		connect.WithInterceptors(h.ValidateInterceptor()),
	))
	mux.Handle("/grpc/auth/", http.StripPrefix("/grpc/auth", api))
	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(userv1connect.UserServiceName),
		compress1KB,
	))
	mux.Handle(grpcreflect.NewHandlerV1(
		grpcreflect.NewStaticReflector(userv1connect.UserServiceName),
		compress1KB,
	))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(
		grpcreflect.NewStaticReflector(userv1connect.UserServiceName),
		compress1KB,
	))
	return h2c.NewHandler(
		mux,
		&http2.Server{})
}

func NewHandlerTest(
	srv UserService,
	logger *zap.Logger,
) *Handler {

	return &Handler{
		svc: srv,
		log: logger,
	}
}
