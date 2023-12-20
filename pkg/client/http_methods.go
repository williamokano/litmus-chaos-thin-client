package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/williamokano/litmus-chaos-thin-client/pkg/utils"
)

type httpResponse[T any] struct {
	data         T
	httpResponse *http.Response
}

func litmusHttpDo[R any](client *http.Client, req *http.Request) (httpResponse[R], error) {
	res, err := client.Do(req)
	if err != nil {
		return httpResponse[R]{
			data:         *new(R),
			httpResponse: nil,
		}, fmt.Errorf("failed to execute http %s request: %v", req.Method, err)
	}

	parsedRes, err := utils.AsJson[LitmusHttpResponse[R]](res)
	if err != nil {
		return httpResponse[R]{
			data:         *new(R),
			httpResponse: nil,
		}, fmt.Errorf("failed to read json body from http %s request: %v", req.Method, err)
	}

	return httpResponse[R]{
		data:         parsedRes.Data,
		httpResponse: res,
	}, nil
}

func litmusGet[R any](client *http.Client, path string) (httpResponse[R], error) {
	var response R
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return httpResponse[R]{
			data:         response,
			httpResponse: nil,
		}, fmt.Errorf("failed to create request: %v", err)
	}

	return litmusHttpDo[R](client, req)
}

func litmusPost[R any](client *http.Client, path string, payload any) (httpResponse[R], error) {
	var response R
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return httpResponse[R]{
			data:         response,
			httpResponse: nil,
		}, fmt.Errorf("failed to marshal payload of type %T: %v", response, err)
	}

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewReader(payloadBytes))
	if err != nil {
		return httpResponse[R]{
			data:         response,
			httpResponse: nil,
		}, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	return litmusHttpDo[R](client, req)
}
