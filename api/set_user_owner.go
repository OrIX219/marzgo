package api

import (
	"encoding/json"
	"fmt"

	"github.com/OrIX219/marzgo/api/models"
)

type SetUserOwner struct {
	AdminUsername string
}

func (SetUserOwner) Body() (RequestBody, error) {
	return nil, nil
}

func (p SetUserOwner) Query() (QueryParams, error) {
	params := make(QueryParams)
	params.AddStringNonEmpty("admin_username", p.AdminUsername)

	return params, nil
}

func (SetUserOwner) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) SetUserOwner(username string, params SetUserOwner) (models.UserResponse, error) {
	var user models.UserResponse
	resp, err := c.Request("PUT", fmt.Sprintf("api/user/%s/set-owner", username), params)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(resp, &user)

	return user, err
}
