package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type LitmusClient struct {
	baseUrl     *url.URL
	httpClient  *http.Client
	credentials LitmusCredentials
}

type LitmusHttpResponse[T any] struct {
	Data T `json:"data"`
}

type LitmusCredentials struct {
	Username string
	Password string
	Token    string
}

// Deprecated: this function will be removed in the future. Please use NewClientFromCredentials instead
func NewLitmusClient(host string, token string) (*LitmusClient, error) {
	baseUrl, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base url: %v", err)
	}

	client := &LitmusClient{
		credentials: LitmusCredentials{
			Token: token,
		},
		baseUrl: baseUrl,
		httpClient: &http.Client{
			Transport: &bearerTransport{
				baseUrl: baseUrl,
				token:   token,
			},
		},
	}

	return client, nil
}

func getTokenFromUserPassLogin(host string, credentials LitmusCredentials) (string, error) {
	if credentials.Username == "" {
		return "", errors.New("missing Username")
	}

	if credentials.Password == "" {
		return "", errors.New("missing Password")
	}

	payload := struct {
		Password string `json:"password"`
		Username string `json:"username"`
	}{
		Password: credentials.Password,
		Username: credentials.Username,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to create payload")
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/auth/login", host), bytes.NewReader(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create /auth/login request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute login request: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("expected status_code 200, got %d", res.StatusCode))
	}

	resBody := struct {
		AccessToken string                 `json:"accessToken"`
		Extra       map[string]interface{} `json:"-"`
	}{}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body bytes: %v", err)
	}
	err = json.Unmarshal(bodyBytes, &resBody)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal login response: %v", err)
	}

	return resBody.AccessToken, nil
}

func NewClientFromCredentials(host string, credentials LitmusCredentials) (*LitmusClient, error) {
	if credentials.Token == "" {
		token, err := getTokenFromUserPassLogin(host, credentials)
		if err != nil {
			return nil, fmt.Errorf("failed to obtain Token from user/pass combination: %v", err)
		}
		credentials.Token = token
	}

	baseUrl, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base url: %v", err)
	}

	client := &LitmusClient{
		credentials: credentials,
		baseUrl:     baseUrl,
		httpClient: &http.Client{
			Transport: &bearerTransport{
				baseUrl: baseUrl,
				token:   credentials.Token,
			},
		},
	}

	return client, nil
}
