package account

import (
	"context"
	"net/http"
	"time"

	"google.golang.org/grpc"

	pb "code.ysitd.cloud/api/account"
)

type Client interface {
	Close()
	GetTransport() string
	ValidateUserPassword(ctx context.Context, username, password string) (*pb.ValidateUserReply, error)
	GetUserInfo(ctx context.Context, username string) (*pb.GetUserInfoReply, error)
	GetTokenInfo(ctx context.Context, token string) (*pb.GetTokenInfoReply, error)
}

func NewClient(transport, endpoint string) (Client, error) {
	if transport == "grpc" {
		return NewGrpcClient(endpoint)
	}

	return NewHTTPClient(transport, endpoint)
}

func NewGrpcClient(endpoint string) (*GrpcClient, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewAccountClient(conn)

	return &GrpcClient{
		Open:   true,
		Conn:   conn,
		Client: client,
	}, nil
}

func NewHTTPClient(transport, endpoint string) (*HttpClient, error) {
	return &HttpClient{
		Transport: transport,
		Endpoint:  endpoint,
		Client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}, nil
}

func NewGatewayClient(endpoint, token string) *GatewayClient {
	return &GatewayClient{
		Endpoint: endpoint,
		Token:    token,
		Client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}
