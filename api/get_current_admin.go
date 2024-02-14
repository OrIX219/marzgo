package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

func (c *Client) GetCurrentAdmin() (models.Admin, error) {
	var admin models.Admin
	resp, err := c.MakeRequest("GET", "api/admin/", ContentTypeNone, nil, nil)
	if err != nil {
		return admin, err
	}

	err = json.Unmarshal(resp, &admin)

	return admin, err
}
