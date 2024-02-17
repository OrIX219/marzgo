package api

import (
	"encoding/json"
	"fmt"

	"github.com/OrIX219/marzgo/api/models"
)

type ModifyAdmin models.AdminModify

func (c ModifyAdmin) Body() (RequestBody, error) {
	return models.AdminModify(c), nil
}

func (ModifyAdmin) Query() (QueryParams, error) {
	return nil, nil
}

func (ModifyAdmin) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) ModifyAdmin(username string, params ModifyAdmin) (models.Admin, error) {
	var admin models.Admin
	resp, err := c.Request("PUT", fmt.Sprintf("api/admin/%s", username), params)
	if err != nil {
		return admin, err
	}

	err = json.Unmarshal(resp, &admin)

	return admin, err
}
