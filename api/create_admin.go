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

func (c *Client) CreateAdmin(username, password string) (models.Admin, error) {
	var admin models.Admin
	resp, err := c.Request("POST", "api/admin",
		CreateAdmin{
			Username: username,
			Password: password,
		})
	if err != nil {
		return admin, err
	}

	err = json.Unmarshal(resp, &admin)

	return admin, err
}
