package userclient

import (
	"context"
	"net"
	"net/http"
	"net/url"

	userv1 "github.com/Astemirdum/transactions/proto/user/v1"
	"github.com/Astemirdum/transactions/proto/user/v1/userv1connect"
	"github.com/Astemirdum/transactions/tx-balance/config"
	"github.com/bufbuild/connect-go"
)

type AuthClient struct {
	client Auth // userv1connect.UserServiceClient
}

func NewAuthClient(cfg config.AuthClient) *AuthClient {
	const (
		api    = "/grpc/auth"
		scheme = "http"
	)
	u := url.URL{
		Scheme: scheme,
		Host:   net.JoinHostPort(cfg.Host, cfg.Port),
		Path:   api,
	}
	client := userv1connect.NewUserServiceClient(
		&http.Client{Timeout: cfg.Timeout},
		u.String(),
		connect.WithGRPC(),
		connect.WithGRPCWeb(),
	)
	return &AuthClient{
		client: client,
	}
}

type Auth interface {
	Auth(context.Context, *connect.Request[userv1.AuthRequest],
	) (*connect.Response[userv1.AuthResponse], error)
}

func (c *AuthClient) Auth(ctx context.Context, session string) (int32, error) {
	req := connect.NewRequest(&userv1.AuthRequest{
		SessionId: &userv1.SessionID{Id: session},
	})
	resp, err := c.client.Auth(ctx, req)
	if err != nil {
		return 0, err
	}
	return resp.Msg.UserId, nil
}
