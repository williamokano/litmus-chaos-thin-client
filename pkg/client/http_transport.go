package client

import (
	"fmt"
	"net/http"
	"net/url"
)

type bearerTransport struct {
	baseUrl *url.URL
	token   string
}

func (t *bearerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Inject the default stuff
	if req.URL.Host == "" {
		req.URL.Host = t.baseUrl.Host
		req.URL.Scheme = t.baseUrl.Scheme
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.token))
	return http.DefaultTransport.RoundTrip(req)
}
