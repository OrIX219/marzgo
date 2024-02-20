package api

import (
	"encoding/json"
	"fmt"

	"github.com/OrIX219/marzgo/api/models"
)

type GetUserTemplate struct{}

func (GetUserTemplate) Body() RequestBody {
	return nil
}

func (p GetUserTemplate) Query() QueryParams {
	return nil
}

func (GetUserTemplate) Headers() Headers {
	return nil
}

// GetUserTemplate returns a user template with a specified id.
func (c *Client) GetUserTemplate(id int) (models.UserTemplateResponse, error) {
	var template models.UserTemplateResponse
	resp, err := c.Request("GET", fmt.Sprintf("api/user_template/%d", id), GetUserTemplate{})
	if err != nil {
		return template, err
	}

	err = json.Unmarshal(resp, &template)

	return template, err
}
