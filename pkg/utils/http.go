package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func AsJson[R any](res *http.Response) (R, error) {
	var jsonResponse R
	bodyBytes, err := io.ReadAll(res.Body)
	x := string(bodyBytes)
	_ = x
	if err != nil {
		return jsonResponse, fmt.Errorf("failed to read bytes from body: %v", err)
	}

	err = json.Unmarshal(bodyBytes, &jsonResponse)
	if err != nil {
		return jsonResponse, fmt.Errorf("failed to unmarshall json response of type %T: %v", jsonResponse, err)
	}

	return jsonResponse, nil
}
