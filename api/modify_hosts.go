package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type ModifyHosts models.HostsModify

func (p ModifyHosts) Body() (RequestBody, error) {
	return models.HostsModify(p), nil
}

func (ModifyHosts) Query() (QueryParams, error) {
	return nil, nil
}

func (ModifyHosts) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) ModifyHosts(params ModifyHosts) (models.ProxyHosts, error) {
	var hosts models.ProxyHosts
	resp, err := c.Request("PUT", "api/hosts", params)
	if err != nil {
		return hosts, err
	}

	err = json.Unmarshal(resp, &hosts)

	return hosts, err
}