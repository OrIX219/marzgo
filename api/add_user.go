package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type AddUser models.UserCreate

func (c AddUser) Body() RequestBody {
	return models.UserCreate(c)
}

func (AddUser) Query() QueryParams {
	return nil
}

func (c AddUser) Headers() Headers {
	return nil
}

// AddUser adds a new user with specified params and returns a resulting user.
//
// Required params are Username and Proxies.
func (c *Client) AddUser(params AddUser) (models.UserResponse, error) {
	var user models.UserResponse
	resp, err := c.Request("POST", "api/user", params)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(resp, &user)

	return user, err
}
