package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type CreateAdmin struct {
	Username string
	Password string
	IsSudo   bool
}

func (CreateAdmin) Method() string {
	return "POST"
}

func (CreateAdmin) Endpoint() string {
	return "api/admin"
}

func (c CreateAdmin) Params() (RequestParams, error) {
	params := make(RequestParams)
	params.AddString("username", c.Username)
	params.AddString("password", c.Password)
	params.AddBool("is_sudo", c.IsSudo)

	return params, nil
}

func (c CreateAdmin) Headers() (RequestParams, error) {
	return nil, nil
}

func (c *Client) CreateAdmin(params CreateAdmin) (models.Admin, error) {
	var admin models.Admin
	resp, err := c.Request(params)
	if err != nil {
		return admin, err
	}

	err = json.Unmarshal(resp, &admin)

	return admin, err
}
