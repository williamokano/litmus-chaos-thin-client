package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/williamokano/litmus-chaos-thin-client/internal/utils"
	"github.com/williamokano/litmus-chaos-thin-client/pkg/client"
	"github.com/williamokano/litmus-chaos-thin-client/pkg/entities"
	"github.com/williamokano/litmus-chaos-thin-client/pkg/graphql"
)

type CliFlags struct {
	host          string
	projectId     string
	token         string
	EnvironmentID string
	Name          string
	Tags          utils.StringSlice
	Type          string
	Description   string
}

func main() {
	// Create flags with some defaults
	flags := CliFlags{}

	flag.StringVar(&flags.host, "host", "", "Host")
	flag.StringVar(&flags.projectId, "projectID", "", "ProjectID")
	flag.StringVar(&flags.token, "token", "", "Token")
	flag.StringVar(&flags.EnvironmentID, "ID", "", "Environment ID")
	flag.StringVar(&flags.Name, "name", "", "Name")
	flag.Var(&flags.Tags, "tag", "Tag. Providing multiple times becomes a list")
	flag.StringVar(&flags.Type, "type", "", "Type")
	flag.StringVar(&flags.Description, "description", "", "Description")
	flag.Parse()

	flags.host = utils.StringCoalesce(flags.host, os.Getenv("LITMUS_CHAOS_HOST"))
	flags.token = utils.StringCoalesce(flags.token, os.Getenv("LITMUS_CHAOS_TOKEN"))

	litmusClient, err := client.NewClientFromCredentials(flags.host, client.LitmusCredentials{
		Token: flags.token,
	})
	if err != nil {
		panic(err)
	}
	request := graphql.UpdateEnvironmentRequest{
		EnvironmentID: flags.EnvironmentID,
		Name:          flags.Name,
		Description:   flags.Description,
		Type:          entities.EnvironmentType(flags.Type),
		Tags:          flags.Tags,
	}
	res, err := litmusClient.UpdateEnvironment(flags.projectId, request)
	_ = json.NewEncoder(os.Stdout).Encode(res)
}
