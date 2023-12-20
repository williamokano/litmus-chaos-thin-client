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
