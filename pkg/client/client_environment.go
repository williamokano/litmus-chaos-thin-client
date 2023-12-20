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
		return nil, fmt.Errorf("failed to get environment with ID %s from project ID %s: %v", environmentId, projectId, err)
	}

	return &query.GetEnvironment.Environment, nil
}
