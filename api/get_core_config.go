package api

import (
	"encoding/json"
)

type GetCoreConfig struct{}

func (GetCoreConfig) Body() RequestBody {
	return nil
}

func (GetCoreConfig) Query() QueryParams {
	return nil
}

func (GetCoreConfig) Headers() Headers {
	return nil
}

// GetCoreConfig returns a core config in raw json.
func (c *Client) GetCoreConfig() (json.RawMessage, error) {
	return c.Request("GET", "api/core/config", GetCoreConfig{})
}
