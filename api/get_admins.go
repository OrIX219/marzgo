package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/responses"
)

type GetAdminsParams struct {
	Offset   int
	Limit    int
	Username string
}

func (GetAdminsParams) Method() string {
	return "GET"
}

func (GetAdminsParams) Endpoint() string {
	return "api/admins"
}

func (c GetAdminsParams) Params() (RequestParams, error) {
	params := make(RequestParams)
	params.AddNonZero("offset", c.Offset)
	params.AddNonZero("limit", c.Limit)
	params.AddNonEmpty("username", c.Username)

	return params, nil
}

func (c *Client) GetAdmins(params GetAdminsParams) ([]responses.Admin, error) {
	admins := []responses.Admin{}
	resp, err := c.Request(params)
	if err != nil {
		return admins, err
	}

	err = json.Unmarshal(resp, &admins)

	return admins, err
}
