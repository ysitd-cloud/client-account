package account

import (
	"context"

	api "code.ysitd.cloud/api/account"
)

type GrpcClient struct {
	Open     bool
	Endpoint string
	Client   *api.Client `inject:""`
}

func (c *GrpcClient) GetTransport() string {
	return TransportGrpc
}

func (c *GrpcClient) ValidateUserPassword(ctx context.Context, username, password string) (*api.ValidateUserReply, error) {
	return c.Client.ValidateUserPassword(ctx, &api.ValidateUserRequest{
		Username: username,
		Password: password,
	})
}

func (c *GrpcClient) GetUserInfo(ctx context.Context, username string) (*api.GetUserInfoReply, error) {
	return c.Client.GetUserInfo(ctx, &api.GetUserInfoRequest{
		Username: username,
	})
}

func (c *GrpcClient) GetTokenInfo(ctx context.Context, token string) (*api.GetTokenInfoReply, error) {
	return c.Client.GetTokenInfo(ctx, &api.GetTokenInfoRequest{
		Token: token,
	})
}
