package account

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	pb "code.ysitd.cloud/api/account"
)

type HttpClient struct {
	Transport string
	Endpoint  string
	Client    *http.Client `inject:""`
}

func (c *HttpClient) url(path string) *url.URL {
	var u url.URL
	u.Scheme = c.Transport
	u.Host = c.Endpoint
	u.Path = path
	return &u
}

func (c *HttpClient) Close() {}

func (c *HttpClient) GetTransport() string {
	return c.Transport
}

func (c *HttpClient) ValidateUserPassword(ctx context.Context, username, password string) (*pb.ValidateUserReply, error) {
	body, err := json.Marshal(&pb.ValidateUserRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.url("/validate").String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var reply pb.ValidateUserReply

	replyBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(replyBody, &reply)

	return &reply, err
}

func (c *HttpClient) GetUserInfo(ctx context.Context, username string) (*pb.GetUserInfoReply, error) {
	req, err := http.NewRequest("GET", c.url("/user/"+username).String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var reply pb.GetUserInfoReply

	replyBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(replyBody, &reply)

	return &reply, err
}

func (c *HttpClient) GetTokenInfo(ctx context.Context, token string) (*pb.GetTokenInfoReply, error) {
	req, err := http.NewRequest("GET", c.url("/token/"+token).String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var reply pb.GetTokenInfoReply

	replyBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(replyBody, &reply)

	return &reply, err
}
