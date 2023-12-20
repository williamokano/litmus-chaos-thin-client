package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/williamokano/litmus-chaos-thin-client/internal/utils"
	"github.com/williamokano/litmus-chaos-thin-client/pkg/client"
	"github.com/williamokano/litmus-chaos-thin-client/pkg/graphql"
)

type GetProjectInput struct {
	host      string
	projectId string
	token     string
}

func main() {
	// Create input with some defaults
	input := GetProjectInput{}

	flag.StringVar(&input.host, "host", "", "Host")
	flag.StringVar(&input.projectId, "projectID", "", "ProjectID")
	flag.StringVar(&input.token, "token", "", "Token")
	flag.Parse()

	input.host = utils.StringCoalesce(input.host, os.Getenv("LITMUS_CHAOS_HOST"))
	input.token = utils.StringCoalesce(input.token, os.Getenv("LITMUS_CHAOS_TOKEN"))

	litmusClient, err := client.NewClientFromCredentials(input.host, client.LitmusCredentials{
		Token: input.token,
	})
	if err != nil {
		panic(err)
	}

	request := graphql.ListEnvironmentRequest{
		Pagination: &graphql.Pagination{
			Page:  0,
			Limit: 30,
		},
	}
	res, err := litmusClient.ListEnvironments(input.projectId, request)
	_ = json.NewEncoder(os.Stdout).Encode(res)
}
