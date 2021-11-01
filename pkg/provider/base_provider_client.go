package provider

import (
	"errors"
	"fmt"

	"github.com/Alexamakans/wharf-common-api-client/pkg/apiclient"
	"github.com/iver-wharf/wharf-core/pkg/problem"
)

type ProjectIdentifier = apiclient.ProjectIdentifier
type WharfBranch = apiclient.WharfBranch
type WharfProject = apiclient.WharfProject

type Client struct {
	apiclient.BaseClient
	ProviderURL string
}

func (c *Client) FetchFile(pi ProjectIdentifier, fileName string) ([]byte, error) {
	return []byte{}, nil
}

func (c *Client) FetchBranches(pi ProjectIdentifier) ([]WharfBranch, error) {
	return []WharfBranch{}, nil
}

func (c *Client) FetchProjectByGroupAndProjectName(groupName, projectName string) (WharfProject, error) {
	pi := ProjectIdentifier{
		Values: []string{groupName, projectName},
	}
	path := fmt.Sprintf("api/project/%s", pi.ToPathEscapedString())
	var project WharfProject
	err := apiclient.DoGetUnmarshal(&project, c, c.ProviderURL, path)
	if err != nil {
		prob := problem.Response{}
		if ok := errors.As(err, &prob); ok {
			return WharfProject{}, fmt.Errorf("%s: %w", prob.Detail, prob)
		}
		return WharfProject{}, fmt.Errorf("failed getting project named %s in %s: %w", projectName, groupName, err)
	}

	return project, nil
}
