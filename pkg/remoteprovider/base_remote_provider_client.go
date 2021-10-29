package remoteprovider

import "github.com/iver-wharf/wharf-api-communication/pkg/apiclient"

type ProjectIdentifier = apiclient.ProjectIdentifier
type WharfBranch = apiclient.WharfBranch
type WharfProject = apiclient.WharfProject

type Client interface {
	apiclient.Client

	// WharfProjectToIdentifier extracts the required information to uniquely identify
	// a project when communicating with the API, and returns it in the form
	// of a ProjectIdentifier object.
	//
	// The order of the values are up to the implementation.
	WharfProjectToIdentifier(project apiclient.WharfProject) apiclient.ProjectIdentifier
}

// Client is to be used for communication between a Wharf provider and
// a remote provider, like AzureDevOps, GitHub, or GitLab.
//
// Implements remoteprovider.Client.
type BaseClient struct {
	apiclient.BaseClient
	RemoteProviderURL string
}
