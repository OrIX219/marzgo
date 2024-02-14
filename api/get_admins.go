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

func (c GetAdmins) Params() (RequestParams, error) {
	params := make(RequestParams)
	params.AddIntNonZero("offset", c.Offset)
	params.AddIntNonZero("limit", c.Limit)
	params.AddStringNonEmpty("username", c.Username)

	return params, nil
}

func (GetAdmins) Headers() (RequestParams, error) {
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
