package provider

import (
	"errors"
	"fmt"

	"github.com/Alexamakans/wharf-api-communication/pkg/apiclient"
	"github.com/iver-wharf/wharf-core/pkg/problem"
)

type Client struct {
	apiclient.BaseClient
	ProviderURL string
}

func (c *Client) FetchFile(pi apiclient.ProjectIdentifier, fileName string) ([]byte, error) {
	return []byte{}, nil
}

func (c *Client) FetchBranches(pi apiclient.ProjectIdentifier) ([]apiclient.WharfBranch, error) {
	return []apiclient.WharfBranch{}, nil
}

func (c *Client) FetchProjectByGroupAndProjectName(groupName, projectName string) (apiclient.WharfProject, error) {
	pi := apiclient.ProjectIdentifier{
		Values: []string{groupName, projectName},
	}
	path := fmt.Sprintf("api/project/%s", pi.ToPathEscapedString())
	var project apiclient.WharfProject
	err := apiclient.DoGetUnmarshal(&project, c, c.ProviderURL, path)
	if err != nil {
		prob := problem.Response{}
		if ok := errors.As(err, &prob); ok {
			return apiclient.WharfProject{}, fmt.Errorf("%s: %w", prob.Detail, prob)
		}
		return apiclient.WharfProject{}, fmt.Errorf("failed getting project named %s in %s: %w", projectName, groupName, err)
	}

	return project, nil
}
