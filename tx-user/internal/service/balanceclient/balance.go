package balanceclient

import (
	"context"
	"net"
	"net/http"
	"net/url"

	balancev1 "github.com/Astemirdum/transactions/proto/balance/v1"
	"github.com/Astemirdum/transactions/proto/balance/v1/balancev1connect"
	"github.com/Astemirdum/transactions/tx-user/config"
	"github.com/bufbuild/connect-go"
)

type BalanceClient struct {
	client Balance // balancev1connect.BalanceServiceClient
}

func NewBalanceClient(cfg config.BalanceClient) *BalanceClient {
	const (
		api    = "/grpc/balance"
		scheme = "http"
	)
	u := url.URL{
		Scheme: scheme,
		Host:   net.JoinHostPort(cfg.Host, cfg.Port),
		Path:   api,
	}
	client := balancev1connect.NewBalanceServiceClient(
		&http.Client{Timeout: cfg.Timeout},
		u.String(),
		connect.WithGRPC(),
		connect.WithGRPCWeb(),
	)
	return &BalanceClient{
		client: client,
	}
}

type Balance interface {
	CreateBalance(context.Context, *connect.Request[balancev1.CreateBalanceRequest],
	) (*connect.Response[balancev1.CreateBalanceResponse], error)
}

func (c *BalanceClient) CreateBalance(ctx context.Context, userID int) (int64, error) {
	req := connect.NewRequest(&balancev1.CreateBalanceRequest{
		UserId: int32(userID),
	})
	resp, err := c.client.CreateBalance(ctx, req)
	if err != nil {
		return 0, err
	}
	return resp.Msg.Cash, nil
}
