package api

import (
	"encoding/json"
	"fmt"

	"github.com/OrIX219/marzgo/api/models"
)

type ModifyUserTemplate models.UserTemplateModify

func (c ModifyUserTemplate) Body() (RequestBody, error) {
	return models.UserTemplateModify(c), nil
}

func (ModifyUserTemplate) Query() (QueryParams, error) {
	return nil, nil
}

func (c ModifyUserTemplate) Headers() (Headers, error) {
	return nil, nil
}

// ModifyUserTemplate modifies a user template with a specified id and returns a resulting template.
func (c *Client) ModifyUserTemplate(id int, params ModifyUserTemplate) (models.UserTemplateResponse, error) {
	var user models.UserTemplateResponse
	resp, err := c.Request("PUT", fmt.Sprintf("api/user_template/%d", id), params)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(resp, &user)

	return user, err
}
