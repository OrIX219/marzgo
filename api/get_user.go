package api

import (
	"encoding/json"
	"fmt"

	"github.com/OrIX219/marzgo/api/models"
)

type GetUser struct{}

func (GetUser) Body() (RequestBody, error) {
	return nil, nil
}

func (GetUser) Query() (QueryParams, error) {
	return nil, nil
}

func (GetUser) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) GetUser(username string) (models.UserResponse, error) {
	var user models.UserResponse
	resp, err := c.Request("GET", fmt.Sprintf("api/user/%s", username), GetUser{})
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(resp, &user)

	return user, err
}
