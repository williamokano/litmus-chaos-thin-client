package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/williamokano/litmus-chaos-thin-client/pkg/client"
)

type GetProjectInput struct {
	host      string
	projectId string
	token     string
}

func main() {
	input := GetProjectInput{}
	flag.StringVar(&input.host, "host", "", "Host")
	flag.StringVar(&input.projectId, "projectId", "", "ProjectID")
	flag.StringVar(&input.token, "token", "", "Token")
	flag.Parse()

	litmusClient, err := client.NewClientFromCredentials(input.host, client.LitmusCredentials{
		Token: input.token,
	})
	if err != nil {
		panic(err)
	}

	res, err := litmusClient.GetProjectById(input.projectId)

	_, _ = fmt.Fprintf(os.Stdout, "%+v", res)
}
