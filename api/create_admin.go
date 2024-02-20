package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type CreateAdmin models.AdminCreate

func (c CreateAdmin) Body() (RequestBody, error) {
	return models.AdminCreate(c), nil
}

func (CreateAdmin) Query() (QueryParams, error) {
	return nil, nil
}

func (c CreateAdmin) Headers() (Headers, error) {
	return nil, nil
}

// CreateAdmin creates a new admin with specified params and
// returns resulting admin.
//
// Required params are Username, Password and IsSudo (default false).
func (c *Client) CreateAdmin(params CreateAdmin) (models.Admin, error) {
	var admin models.Admin
	resp, err := c.Request("POST", "api/admin", params)
	if err != nil {
		return admin, err
	}

	err = json.Unmarshal(resp, &admin)

	return admin, err
}
