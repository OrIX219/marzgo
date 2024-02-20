package api

import (
	"encoding/json"
	"fmt"

	"github.com/OrIX219/marzgo/api/models"
)

type ModifyUserTemplate models.UserTemplateModify

func (c ModifyUserTemplate) Body() RequestBody {
	return models.UserTemplateModify(c)
}

func (ModifyUserTemplate) Query() QueryParams {
	return nil
}

func (c ModifyUserTemplate) Headers() Headers {
	return nil
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
