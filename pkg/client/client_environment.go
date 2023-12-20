package client

import (
	"fmt"

	hasuragraphql "github.com/hasura/go-graphql-client"
	"github.com/williamokano/litmus-chaos-thin-client/pkg/entities"
	"github.com/williamokano/litmus-chaos-thin-client/pkg/graphql"
)

func (c *LitmusClient) GetEnvironmentByID(projectId string, environmentId string) (*entities.Environment, error) {
	query := graphql.GetEnvironmentQuery{}
	args := map[string]interface{}{
		"projectID":     *hasuragraphql.NewID(projectId),
		"environmentID": *hasuragraphql.NewID(environmentId),
	}
	err := c.graphqlClient.Query(c.ctx, &query, args)

	if err != nil {
		return nil, fmt.Errorf("failed to get environment with ID %s from project ID %s: %w", environmentId, projectId, err)
	}

	return &query.GetEnvironment.Environment, nil
}

func (c *LitmusClient) ListEnvironments(projectId string, request graphql.ListEnvironmentRequest) ([]*entities.Environment, error) {
	query := graphql.ListEnvironmentsQuery{}
	args := map[string]interface{}{
		"projectID": *hasuragraphql.NewID(projectId),
		"request":   request,
	}
	err := c.graphqlClient.Query(c.ctx, &query, args)

	if err != nil {
		return nil, fmt.Errorf("failed to list environments from project ID %s: %w", projectId, err)
	}

	return query.ListEnvironments.Environments, nil
}

func (c *LitmusClient) CreateEnvironment(projectId string, request graphql.CreateEnvironmentRequest) (*entities.Environment, error) {
	// Check if ID is provided, otherwise fallback to default implementation
	if request.EnvironmentID == "" {
		request.EnvironmentID = request.IDFromName()
	}

	mutation := graphql.CreateEnvironmentMutation{}
	args := map[string]interface{}{
		"projectID": *hasuragraphql.NewID(projectId),
		"request":   request,
	}
	err := c.graphqlClient.Mutate(c.ctx, &mutation, args)

	if err != nil {
		return nil, fmt.Errorf("failed to create environment on project ID %s: %w", projectId, err)
	}

	return &mutation.CreateEnvironment.Environment, nil
}

func (c *LitmusClient) UpdateEnvironment(projectId string, request graphql.UpdateEnvironmentRequest) (string, error) {
	mutation := graphql.UpdateEnvironmentMutation{}
	args := map[string]interface{}{
		"projectID": *hasuragraphql.NewID(projectId),
		"request":   request,
	}
	err := c.graphqlClient.Mutate(c.ctx, &mutation, args)

	if err != nil {
		return "", fmt.Errorf("failed to update environment on project ID %s: %w", projectId, err)
	}

	return mutation.UpdateEnvironment, nil
}
