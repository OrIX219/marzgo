package api

import (
	"encoding/json"
	"fmt"

	"github.com/OrIX219/marzgo/api/models"
)

type ModifyAdmin struct {
	Username string
	Password string
	IsSudo   bool
}

func (ModifyAdmin) Method() string {
	return "PUT"
}

func (p ModifyAdmin) Endpoint() string {
	return fmt.Sprintf("api/admin/%s", p.Username)
}

func (c ModifyAdmin) Body() (BodyParams, error) {
	params := make(BodyParams)
	params.AddStringNonEmpty("password", c.Password)
	params.AddBool("is_sudo", c.IsSudo)

	return params, nil
}

func (ModifyAdmin) Url() (UrlParams, error) {
	return nil, nil
}

func (ModifyAdmin) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) ModifyAdmin(params ModifyAdmin) (models.Admin, error) {
	admin := models.Admin{}
	resp, err := c.Request(params)
	if err != nil {
		return admin, err
	}

	err = json.Unmarshal(resp, &admin)

	return admin, err
}
