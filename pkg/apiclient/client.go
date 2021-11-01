package apiclient

import "context"

type Client interface {
	// GetContext gets the context for the client.
	//
	// This method is implemented by the BaseClient type for convenience.
	GetContext() context.Context
	// GetToken gets the token that is used when sending requests to the API.
	//
	// This method is implemented by the BaseClient type for convenience.
	GetToken() string
	// SetContext sets the context to use for the connection.
	//
	// This method is implemented by the BaseClient type for convenience.
	SetContext(ctx context.Context)
	// SetToken sets the token that is used when sending requests to the remote provider.
	//
	// This method is implemented by the BaseClient type for convenience.
	SetToken(token string)
}
