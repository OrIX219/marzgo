package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type ModifyHosts models.HostsModify

func (p ModifyHosts) Body() RequestBody {
	return models.HostsModify(p)
}

func (ModifyHosts) Query() QueryParams {
	return nil
}

func (ModifyHosts) Headers() Headers {
	return nil
}

// ModifyHosts modifies hosts for each specified inbound and returns resulting hosts
// divided in inbounds.
//
// WARNING: If a specified inbound contains an unspecified host
// then that host will be removed from that inbound.
// So to preserve inbound hosts which you don't want to modify it's recommended to
// get all hosts with GetHosts first, modify those and then pass to this function
// (ProxyHosts can be cast to ModifyHosts).
func (c *Client) ModifyHosts(params ModifyHosts) (models.ProxyHosts, error) {
	var hosts models.ProxyHosts
	resp, err := c.Request("PUT", "api/hosts", params)
	if err != nil {
		return hosts, err
	}

	err = json.Unmarshal(resp, &hosts)

	return hosts, err
}
