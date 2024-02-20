package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type GetAdmins struct {
	Offset   int
	Limit    int
	Username string
}

func (GetAdmins) Body() RequestBody {
	return nil
}

func (p GetAdmins) Query() QueryParams {
	params := make(QueryParams)
	params.AddIntNonZero("offset", p.Offset)
	params.AddIntNonZero("limit", p.Limit)
	params.AddStringNonEmpty("username", p.Username)

	return params
}

func (GetAdmins) Headers() Headers {
	return nil
}

// GetAdmins returns all admins offseted and limited by params.
// If Username is specified returns only matching admin if exists.
func (c *Client) GetAdmins(params GetAdmins) ([]models.Admin, error) {
	admins := []models.Admin{}
	resp, err := c.Request("GET", "api/admins", params)
	if err != nil {
		return admins, err
	}

	err = json.Unmarshal(resp, &admins)

	return admins, err
}
