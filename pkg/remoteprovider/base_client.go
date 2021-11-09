package remoteprovider

import (
	"context"

	"github.com/Alexamakans/wharf-common-api-client/pkg/apiclient"
)

// BaseClient implements a subset of the Client interface, and has
// fields that are common to all current implementations of it.
type BaseClient struct {
	apiclient.BaseClient
	RemoteProviderURL string
}

func NewClient(ctx context.Context, token, apiURLPrefix, remoteProviderURL string) *BaseClient {
	return &BaseClient{
		BaseClient:        *apiclient.NewClient(ctx, token, apiURLPrefix),
		RemoteProviderURL: remoteProviderURL,
	}
}

func (c *BaseClient) SetRemoteProviderURL(remoteProviderURL string) {
	c.RemoteProviderURL = remoteProviderURL
}

func (c *BaseClient) GetRemoteProviderURL() string {
	return c.RemoteProviderURL
}
