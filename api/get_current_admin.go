package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/responses"
)

func (c *Client) GetCurrentAdmin() (responses.Admin, error) {
	var admin responses.Admin
	resp, err := c.MakeRequest("GET", "api/admin/", ContentTypeNone, nil)
	if err != nil {
		return admin, err
	}

	err = json.Unmarshal(resp, &admin)

	return admin, err
}
