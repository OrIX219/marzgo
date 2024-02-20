package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type CreateAdmin models.AdminCreate

func (c CreateAdmin) Body() RequestBody {
	return models.AdminCreate(c)
}

func (CreateAdmin) Query() QueryParams {
	return nil
}

func (c CreateAdmin) Headers() Headers {
	return nil
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
