package api

import (
	"encoding/json"
	"fmt"

	"github.com/OrIX219/marzgo/api/models"
)

type ResetUserDataUsage struct{}

func (c ResetUserDataUsage) Body() (RequestBody, error) {
	return nil, nil
}

func (ResetUserDataUsage) Query() (QueryParams, error) {
	return nil, nil
}

func (c ResetUserDataUsage) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) ResetUserDataUsage(username string) (models.UserResponse, error) {
	var user models.UserResponse
	resp, err := c.Request("POST", fmt.Sprintf("api/user/%s/reset", username), ResetUserDataUsage{})
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(resp, &user)

	return user, err
}