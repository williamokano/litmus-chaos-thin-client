package client

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/williamokano/litmus-chaos-thin-client/pkg/entities"
)

type LitmusClient struct {
	baseUrl    *url.URL
	httpClient *http.Client
}

type LitmusHttpResponse[T any] struct {
	Data T `json:"data"`
}

func NewLitmusClient(host string, token string) (*LitmusClient, error) {
	baseUrl, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base url: %v", err)
	}

	client := &LitmusClient{
		httpClient: &http.Client{
			Transport: &bearerTransport{
				baseUrl: baseUrl,
				token:   token,
			},
		},
	}

	return client, nil
}

func (c *LitmusClient) GetProjectById(projectId string) (*entities.Project, error) {
	return litmusGet[*entities.Project](c.httpClient, fmt.Sprintf("/auth/get_project/%s", projectId))
}

func (c *LitmusClient) CreateProject(projectName string) (*entities.Project, error) {
	return litmusPost[*entities.Project](c.httpClient, "/auth/create_project", entities.CreateProjectInput{
		ProjectName: projectName,
	})
}
