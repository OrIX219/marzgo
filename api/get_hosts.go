package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type GetHosts struct{}

func (GetHosts) Body() (RequestBody, error) {
	return nil, nil
}

func (GetHosts) Query() (QueryParams, error) {
	return nil, nil
}

func (GetHosts) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) GetHosts() (models.ProxyHosts, error) {
	var hosts models.ProxyHosts
	resp, err := c.Request("GET", "api/hosts", GetHosts{})
	if err != nil {
		return hosts, err
	}

	err = json.Unmarshal(resp, &hosts)

	return hosts, err
}
