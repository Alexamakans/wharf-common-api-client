package apiclient

import "github.com/iver-wharf/wharf-api/pkg/model/database"

// WharfProject is an alias for wharf-api's database.Project.
// type WharfProject = database.Project

// WharfProject is an alias for wharf-api's database.Project.
// Currently using a hack to pretend that the RemoteProjectID
// addition has been merged.
type WharfProject struct {
	database.Project
	RemoteProjectID string
}

// WharfBranch is an alias for wharf-api's database.Branch.
type WharfBranch = database.Branch
