package remoteprovider

import "github.com/Alexamakans/wharf-common-api-client/pkg/apiclient"

type Client interface {
	apiclient.Client

	// FetchFile fetches a file from the API.
	FetchFile(projectIdentifiers ProjectIdentifier, fileName string) ([]byte, error)
	// FetchBranches fetches branches for a specific project from the API.
	FetchBranches(projectIdentifier ProjectIdentifier) ([]WharfBranch, error)
	// FetchProjectByGroupAndProjectName should only be used if the remote
	// project ID for the project is not available. If it is available then
	// FetchProjectByID should be used instead.
	//
	// Mainly used for migrations for projects added prior to when we started
	// tracking projects by their remote project ID in wharf-api.
	FetchProjectByGroupAndProjectName(groupName, projectName string) (WharfProject, error)

	// GetRemoteProviderURL returns the set remote provider URL.
	//
	// Implemented by BaseClient for convenience.
	GetRemoteProviderURL() string

	// SetRemoteProviderURL sets the remote provider URL.
	//
	// Implemented by BaseClient for convenience.
	SetRemoteProviderURL(remoteProviderURL string)

	// WharfProjectToIdentifier extracts the required information to uniquely
	// identify a project when communicating with the API, and returns it in the
	// form of a ProjectIdentifier object.
	//
	// The order of the values are up to the implementation.
	WharfProjectToIdentifier(project WharfProject) ProjectIdentifier
}
