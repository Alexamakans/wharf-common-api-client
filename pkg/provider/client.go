package provider

import "github.com/Alexamakans/wharf-common-api-client/pkg/apiclient"

// Client is an interface meant to generalize interactions with an API.
// Specifically, this is a collection of methods commonly used by the wharf API
// when communicating with any of the wharf providers, and by the wharf providers when
// communicating with the remote provider.
type Client interface {
	apiclient.Client

	// FetchFile fetches a file from the API.
	FetchFile(project WharfProject, fileName string) ([]byte, error)
	// FetchBranches fetches branches for a specific project from the API.
	FetchBranches(project WharfProject) ([]WharfBranch, error)
	// FetchProjectByGroupAndProjectName should only be used if the remote project ID for the
	// project is not available. If it is available then FetchProjectByID should be used instead.
	//
	// Mainly used for migrations for projects added prior to when we started tracking projects by
	// their remote project ID in wharf-api.
	FetchProjectByGroupAndProjectName(groupName, projectName string) (WharfProject, error)
}
