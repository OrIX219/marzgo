package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/responses"
)

type CreateAdminParams struct {
	Username string
	Password string
	IsSudo   bool
}

func (CreateAdminParams) Method() string {
	return "POST"
}

func (CreateAdminParams) Endpoint() string {
	return "api/admin"
}

func (c CreateAdminParams) Params() (RequestParams, error) {
	params := make(RequestParams)
	params.AddString("username", c.Username)
	params.AddString("password", c.Password)
	params.AddBool("is_sudo", c.IsSudo)

	return params, nil
}

func (c *Client) CreateAdmin(params CreateAdminParams) (responses.Admin, error) {
	admin := responses.Admin{}
	resp, err := c.Request(params)
	if err != nil {
		return admin, err
	}

	err = json.Unmarshal(resp, &admin)

	return admin, err
}
