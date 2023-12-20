package client

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/williamokano/litmus-chaos-thin-client/pkg/entities"
)

func (c *LitmusClient) GetProjectById(projectId string) (*entities.Project, error) {
	res, err := litmusGet[*entities.Project](c.httpClient, fmt.Sprintf("/auth/get_project/%s", projectId))

	if err != nil {
		return res.data, fmt.Errorf("failed to get project by id: %v", err)
	}

	if res.httpResponse.StatusCode != http.StatusOK {
		return res.data, fmt.Errorf("failed to get project by ID")
	}

	return res.data, nil
}

func (c *LitmusClient) CreateProject(projectName string) (*entities.Project, error) {
	res, err := litmusPost[*entities.Project](c.httpClient, "/auth/create_project", entities.CreateProjectInput{
		ProjectName: projectName,
	})

	if err != nil {
		return res.data, fmt.Errorf("failed to create project with name %s: %v", projectName, err)
	}

	// Litmus don't return 201 but rather 200
	// https://github.com/litmuschaos/litmus/blob/master/chaoscenter/authentication/api/handlers/rest/project_handler.go#L338
	if res.httpResponse.StatusCode != http.StatusOK {
		return res.data, errors.New(fmt.Sprintf("api responded %d instead of %d", res.httpResponse.StatusCode, http.StatusOK))
	}

	return res.data, nil
}

func (c *LitmusClient) UpdateProjectName(projectId string, projectName string) (*entities.Project, error) {
	res, err := litmusPost[map[string]interface{}](c.httpClient, "/auth/update_project_name", entities.UpdateProjectNameInput{
		ProjectID:   projectId,
		ProjectName: projectName,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to update project name: %v", err)
	}

	if res.httpResponse.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("expected status_code is %d, returned %d", http.StatusOK, res.httpResponse.StatusCode))
	}

	return c.GetProjectById(projectId)
}
