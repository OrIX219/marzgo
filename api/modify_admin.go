package api

import (
	"encoding/json"
	"fmt"

	"github.com/OrIX219/marzgo/api/responses"
)

type ModifyAdminParams struct {
	Username string
	Password string
	IsSudo   bool
}

func (ModifyAdminParams) Method() string {
	return "PUT"
}

func (p ModifyAdminParams) Endpoint() string {
	return fmt.Sprintf("api/admin/%s", p.Username)
}

func (c ModifyAdminParams) Params() (RequestParams, error) {
	params := make(RequestParams)
	params.AddNonEmpty("password", c.Password)
	params.AddBool("is_sudo", c.IsSudo)

	return params, nil
}

func (c *Client) ModifyAdmin(params ModifyAdminParams) (responses.Admin, error) {
	admin := responses.Admin{}
	resp, err := c.Request(params)
	if err != nil {
		return admin, err
	}

	err = json.Unmarshal(resp, &admin)

	return admin, err
}
