package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type AddUserTemplate models.UserTemplateCreate

func (c AddUserTemplate) Body() (RequestBody, error) {
	return models.UserTemplateCreate(c), nil
}

func (AddUserTemplate) Query() (QueryParams, error) {
	return nil, nil
}

func (c AddUserTemplate) Headers() (Headers, error) {
	return nil, nil
}

// AddUserTemplate adds a new user template with specified params and
// returns resulting template.
func (c *Client) AddUserTemplate(params AddUserTemplate) (models.UserTemplateResponse, error) {
	var user models.UserTemplateResponse
	resp, err := c.Request("POST", "api/user_template", params)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(resp, &user)

	return user, err
}
