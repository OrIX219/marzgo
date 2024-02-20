package api

import (
	"encoding/json"
	"fmt"

	"github.com/OrIX219/marzgo/api/models"
)

type ModifyUser models.UserModify

func (c ModifyUser) Body() RequestBody {
	return models.UserModify(c)
}

func (ModifyUser) Query() QueryParams {
	return nil
}

func (c ModifyUser) Headers() Headers {
	return nil
}

// ModifyUser modifies a user with a specified username and returns a resulting user.
func (c *Client) ModifyUser(username string, params ModifyUser) (models.UserResponse, error) {
	var user models.UserResponse
	resp, err := c.Request("PUT", fmt.Sprintf("api/user/%s", username), params)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(resp, &user)

	return user, err
}
