package account

import (
	"context"
	"net/http"
	"time"

	api "code.ysitd.cloud/api/account"
)

type Client interface {
	GetTransport() string
	ValidateUserPassword(ctx context.Context, username, password string) (*api.ValidateUserReply, error)
	GetUserInfo(ctx context.Context, username string) (*api.GetUserInfoReply, error)
	GetTokenInfo(ctx context.Context, token string) (*api.GetTokenInfoReply, error)
}

func NewClient(transport, endpoint string) Client {
	if transport == TransportGrpc {
		return NewGrpcClient(endpoint)
	} else if transport == TransportGateway {
		return NewGatewayClient(endpoint)
	}

	return NewHTTPClient(transport, endpoint)
}

func NewGrpcClient(endpoint string) *GrpcClient {
	return &GrpcClient{
		Open:   true,
		Client: api.NewClient(endpoint),
	}
}

func NewHTTPClient(transport, endpoint string) *HttpClient {
	return &HttpClient{
		Transport: transport,
		Endpoint:  endpoint,
		Client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func NewGatewayClient(endpoint string) *GatewayClient {
	return &GatewayClient{
		Endpoint: endpoint,
		Client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}
