package api

import (
	"encoding/json"
	"fmt"

	"github.com/OrIX219/marzgo/api/models"
)

type ModifyUser models.UserModify

func (c ModifyUser) Body() (RequestBody, error) {
	return models.UserModify(c), nil
}

func (ModifyUser) Query() (QueryParams, error) {
	return nil, nil
}

func (c ModifyUser) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) ModifyUser(username string, params ModifyUser) (models.UserResponse, error) {
	var user models.UserResponse
	resp, err := c.Request("PUT", fmt.Sprintf("api/user/%s", username), params)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(resp, &user)

	return user, err
}
