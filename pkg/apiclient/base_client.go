package apiclient

import "context"

type BaseClient struct {
	Context      context.Context
	Token        string
	APIURLPrefix string
}

func NewClient(ctx context.Context, token, apiURLPrefix string) *BaseClient {
	return &BaseClient{
		APIURLPrefix: apiURLPrefix,
	}
}

func (c *BaseClient) GetContext() context.Context {
	return c.Context
}

func (c *BaseClient) GetToken() string {
	return c.Token
}

func (c *BaseClient) SetContext(ctx context.Context) {
	c.Context = ctx
}

func (c *BaseClient) SetToken(token string) {
	c.Token = token
}
