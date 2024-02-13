package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/responses"
)

func (c *Client) AdminToken() (responses.Token, error) {
	var token responses.Token
	resp, err := c.MakeRequest("POST", "api/admin/token", ContentTypeUrlEncoded,
		RequestParams{
			"username": c.Username,
			"password": c.Password,
		})
	if err != nil {
		return token, err
	}

	err = json.Unmarshal(resp, &token)

	return token, err
}
