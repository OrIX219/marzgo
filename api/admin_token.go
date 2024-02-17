package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

func (c *Client) AdminToken() (models.Token, error) {
	var token models.Token
	resp, err := c.MakeRequest("POST", "api/admin/token", ContentTypeUrlEncoded,
		models.GetToken{
			Username: c.Username,
			Password: c.Password,
		}, nil, nil)
	if err != nil {
		return token, err
	}

	err = json.Unmarshal(resp, &token)

	return token, err
}
