package builder

import (
	"context"
	"gitee.com/candy44/dbcore-sdk/pbfiles"

	"google.golang.org/grpc"
)

type ClientBuilder struct {
	url  string
	opts []grpc.DialOption
}

func NewClientBuilder(url string) *ClientBuilder {
	return &ClientBuilder{url: url}
}
func (c *ClientBuilder) WithOption(opts ...grpc.DialOption) *ClientBuilder {
	c.opts = append(c.opts, opts...)
	return c
}
func (c *ClientBuilder) Build() (pbfiles.DBServiceClient, error) {
	client, err := grpc.DialContext(context.Background(),
		c.url,
		c.opts...,
	)
	if err != nil {
		return nil, err
	}
	return pbfiles.NewDBServiceClient(client), nil

}
