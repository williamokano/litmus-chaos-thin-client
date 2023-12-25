package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/williamokano/litmus-chaos-thin-client/internal/utils"
	"github.com/williamokano/litmus-chaos-thin-client/pkg/client"
)

type GetProjectInput struct {
	host  string
	token string
}

func main() {
	// Create input with some defaults
	input := GetProjectInput{}

	flag.StringVar(&input.host, "host", "", "Host")
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

	res, err := litmusClient.FetchUsers()
	if err != nil {
		panic(err)
	}
	_ = json.NewEncoder(os.Stdout).Encode(res)
}
