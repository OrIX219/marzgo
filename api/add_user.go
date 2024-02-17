package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type AddUser models.UserCreate

func (c AddUser) Body() (RequestBody, error) {
	return models.UserCreate(c), nil
}

func (AddUser) Query() (QueryParams, error) {
	return nil, nil
}

func (c AddUser) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) AddUser(params AddUser) (models.UserResponse, error) {
	var user models.UserResponse
	resp, err := c.Request("POST", "api/user", params)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(resp, &user)

	return user, err
}
