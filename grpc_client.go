package account

import (
	"context"

	pb "code.ysitd.cloud/api/account"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	Open   bool
	Client pb.AccountClient
	Conn   *grpc.ClientConn `inject:""`
}

func (c *GrpcClient) Close() {
	if c.Open {
		c.Conn.Close()
		c.Open = false
	}
}

func (c *GrpcClient) GetTransport() string {
	return Grpc
}

func (c *GrpcClient) ValidateUserPassword(ctx context.Context, username, password string) (*pb.ValidateUserReply, error) {
	return c.Client.ValidateUserPassword(ctx, &pb.ValidateUserRequest{
		Username: username,
		Password: password,
	})
}

func (c *GrpcClient) GetUserInfo(ctx context.Context, username string) (*pb.GetUserInfoReply, error) {
	return c.Client.GetUserInfo(ctx, &pb.GetUserInfoRequest{
		Username: username,
	})
}

func (c *GrpcClient) GetTokenInfo(ctx context.Context, token string) (*pb.GetTokenInfoReply, error) {
	return c.Client.GetTokenInfo(ctx, &pb.GetTokenInfoRequest{
		Token: token,
	})
}
