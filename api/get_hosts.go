package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type GetHosts struct{}

func (GetHosts) Body() RequestBody {
	return nil
}

func (GetHosts) Query() QueryParams {
	return nil
}

func (GetHosts) Headers() Headers {
	return nil
}

// GetHosts returns all hosts divided in inbounds.
// Key is an inbound tag, value is a slice of hosts.
func (c *Client) GetHosts() (models.ProxyHosts, error) {
	var hosts models.ProxyHosts
	resp, err := c.Request("GET", "api/hosts", GetHosts{})
	if err != nil {
		return hosts, err
	}

	err = json.Unmarshal(resp, &hosts)

	return hosts, err
}
