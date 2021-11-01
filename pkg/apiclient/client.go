package apiclient

import (
	"context"
)

// Client is an interface meant to generalize interactions with an API.
// Specifically, this is a collection of methods commonly used by the wharf API
// when communicating with any of the wharf providers, and by the wharf providers when
// communicating with the remote provider.
type Client interface {
	// FetchFile fetches a file from the API.
	FetchFile(projectIdentifiers ProjectIdentifier, fileName string) ([]byte, error)
	// FetchBranches fetches branches for a specific project from the API.
	FetchBranches(projectIdentifier ProjectIdentifier) ([]WharfBranch, error)
	// FetchProjectByGroupAndProjectName should only be used if the remote project ID for the
	// project is not available. If it is available then FetchProjectByID should be used instead.
	//
	// Mainly used for migrations for projects added prior to when we started tracking projects by
	// their remote project ID in wharf-api.
	FetchProjectByGroupAndProjectName(groupName, projectName string) (WharfProject, error)

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
