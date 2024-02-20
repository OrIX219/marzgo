package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type GetUserTemplates struct {
	Offset int
	Limit  int
}

func (GetUserTemplates) Body() RequestBody {
	return nil
}

func (p GetUserTemplates) Query() QueryParams {
	params := make(QueryParams)
	params.AddIntNonZero("offset", p.Offset)
	params.AddIntNonZero("limit", p.Limit)

	return params
}

func (GetUserTemplates) Headers() Headers {
	return nil
}

// GetUserTemplates returns all user templates.
func (c *Client) GetUserTemplates(params GetUserTemplates) ([]models.UserTemplateResponse, error) {
	templates := []models.UserTemplateResponse{}
	resp, err := c.Request("GET", "api/user_template", params)
	if err != nil {
		return templates, err
	}

	err = json.Unmarshal(resp, &templates)

	return templates, err
}
