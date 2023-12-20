package client

import (
	"fmt"

	"github.com/williamokano/litmus-chaos-thin-client/pkg/entities"
	"github.com/williamokano/litmus-chaos-thin-client/pkg/graphql"
)

func (c *LitmusClient) GetEnvironmentByID(projectId string, environmentId string) (*entities.Environment, error) {
	query := graphql.GetEnvironmentQuery{}

	err := c.graphqlClient.Query(c.ctx, &query, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to get environment with ID %s from project ID %s: %v", environmentId, projectId, err)
	}

	return &query.GetEnvironment.Environment, nil
}
