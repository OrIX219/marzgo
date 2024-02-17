package api

import (
	"encoding/json"
)

type GetCoreConfig struct{}

func (GetCoreConfig) Body() (RequestBody, error) {
	return nil, nil
}

func (GetCoreConfig) Query() (QueryParams, error) {
	return nil, nil
}

func (GetCoreConfig) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) GetCoreConfig() (json.RawMessage, error) {
	return c.Request("GET", "api/core/config", GetCoreConfig{})
}
