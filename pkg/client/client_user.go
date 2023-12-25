package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/williamokano/litmus-chaos-thin-client/pkg/entities"
)

// The abstraction failed there as it doesn't return a data field
func (c *LitmusClient) FetchUsers() ([]entities.User, error) {
	req, err := http.NewRequest(http.MethodGet, "/auth/users", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusUnauthorized {
			return nil, errors.New("you don't have permission to fetch users")
		}

		return nil, errors.New(fmt.Sprintf("failed to fetch users with response code %d", res.StatusCode))
	}
	var users []entities.User
	err = json.NewDecoder(res.Body).Decode(&users)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return users, nil
}

func (c *LitmusClient) FindUserByUsername(username string) (*entities.User, error) {
	users, err := c.FetchUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}

	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("couldn't find user with username %s", username))
}
