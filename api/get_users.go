package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type GetUsers struct {
	Offset   int
	Limit    int
	Username []string
	Status   models.UserStatus
	Sort     string
}

func (GetUsers) Body() (RequestBody, error) {
	return nil, nil
}

func (p GetUsers) Query() (QueryParams, error) {
	params := make(QueryParams)
	params.AddIntNonZero("offset", p.Offset)
	params.AddIntNonZero("limit", p.Limit)
	params.AddStringSliceNonEmpty("username", p.Username)
	params.AddStringNonEmpty("status", string(p.Status))
	params.AddStringNonEmpty("sort", p.Sort)

	return params, nil
}

func (GetUsers) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) GetUsers(params GetUsers) (models.UsersResponse, error) {
	var users models.UsersResponse
	resp, err := c.Request("GET", "api/users", params)
	if err != nil {
		return users, err
	}

	err = json.Unmarshal(resp, &users)

	return users, err
}
