package provider

import (
	"errors"
	"fmt"

	"github.com/Alexamakans/wharf-common-api-client/pkg/apiclient"
	"github.com/iver-wharf/wharf-core/pkg/problem"
)

type BaseClient struct {
	apiclient.BaseClient
	ProviderURL string
}

func (c *BaseClient) FetchFile(project WharfProject, fileName string) ([]byte, error) {
	return apiclient.DoPostBytes(c, c.ProviderURL, "api/project/file", &project, "filename", fileName)
}

func (c *BaseClient) FetchBranches(project WharfProject) ([]WharfBranch, error) {
	var branches []WharfBranch
	if err := apiclient.DoPostUnmarshal(&branches, c, c.ProviderURL, "api/project/branch", &project); err != nil {
		return []WharfBranch{}, err
	}

	return branches, nil
}

func (c *BaseClient) FetchProjectByGroupAndProjectName(groupName, projectName string) (WharfProject, error) {
	var project WharfProject
	err := apiclient.DoPostUnmarshal(&project, c, c.ProviderURL, "api/project", WharfProject{
		GroupName: groupName,
		Name:      projectName,
	})
	if err != nil {
		prob := problem.Response{}
		if ok := errors.As(err, &prob); ok {
			return WharfProject{}, fmt.Errorf("%s: %w", prob.Detail, prob)
		}
		return WharfProject{}, fmt.Errorf("failed getting project named %s in %s: %w", projectName, groupName, err)
	}

	return project, nil
}
