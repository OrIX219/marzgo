package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type GetNodeSettings struct{}

func (GetNodeSettings) Body() (RequestBody, error) {
	return nil, nil
}

func (GetNodeSettings) Query() (QueryParams, error) {
	return nil, nil
}

func (GetNodeSettings) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) GetNodeSettings() (models.NodeSettings, error) {
	var settings models.NodeSettings
	resp, err := c.Request("GET", "api/node/settings", GetNodeSettings{})
	if err != nil {
		return settings, err
	}

	err = json.Unmarshal(resp, &settings)

	return settings, err
}
