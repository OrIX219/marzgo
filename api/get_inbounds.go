package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type GetInbounds struct{}

func (GetInbounds) Body() RequestBody {
	return nil
}

func (GetInbounds) Query() QueryParams {
	return nil
}

func (GetInbounds) Headers() Headers {
	return nil
}

// GetInbounds returns all inbounds divided in protocols.
func (c *Client) GetInbounds() (models.ProxyInbounds, error) {
	var inbounds models.ProxyInbounds
	resp, err := c.Request("GET", "api/inbounds", GetInbounds{})
	if err != nil {
		return inbounds, err
	}

	err = json.Unmarshal(resp, &inbounds)

	return inbounds, err
}
