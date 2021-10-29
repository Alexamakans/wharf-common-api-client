package apiclient

import (
	"context"
)

// BaseClient implements a subset of the Client interface, and has
// fields that are common to all current implementations of it.
type BaseClient struct {
	Context context.Context
	Token   string
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
