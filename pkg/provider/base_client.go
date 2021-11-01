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
	project = stripProject(project)

	return apiclient.DoPostBytes(c, c.ProviderURL, "api/project/file", &project, "filename", fileName)
}

func (c *BaseClient) FetchBranches(project WharfProject) ([]WharfBranch, error) {
	project = stripProject(project)

	var branches []WharfBranch
	if err := apiclient.DoPostUnmarshal(&branches, c, c.ProviderURL, "api/project/branch", &project); err != nil {
		return []WharfBranch{}, err
	}

	return branches, nil
}

func (c *BaseClient) FetchProjectByGroupAndProjectName(project WharfProject) (WharfProject, error) {
	var resProject WharfProject
	err := apiclient.DoPostUnmarshal(&resProject, c, c.ProviderURL, "api/project", stripProject(project))
	if err != nil {
		prob := problem.Response{}
		if ok := errors.As(err, &prob); ok {
			return WharfProject{}, fmt.Errorf("%s: %w", prob.Detail, prob)
		}
		return WharfProject{}, fmt.Errorf("failed getting project named %s in %s: %w", project.Name, project.GroupName, err)
	}

	return resProject, nil
}
