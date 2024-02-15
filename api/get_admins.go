package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type GetAdmins struct {
	Offset   int
	Limit    int
	Username string
}

func (GetAdmins) Method() string {
	return "GET"
}

func (GetAdmins) Endpoint() string {
	return "api/admins"
}

func (GetAdmins) Body() (BodyParams, error) {
	return nil, nil
}

func (p GetAdmins) Url() (UrlParams, error) {
	params := make(UrlParams)
	params.AddIntNonZero("offset", p.Offset)
	params.AddIntNonZero("limit", p.Limit)
	params.AddStringNonEmpty("username", p.Username)

	return params, nil
}

func (GetAdmins) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) GetAdmins(params GetAdmins) ([]models.Admin, error) {
	admins := []models.Admin{}
	resp, err := c.Request(params)
	if err != nil {
		return admins, err
	}

	err = json.Unmarshal(resp, &admins)

	return admins, err
}
