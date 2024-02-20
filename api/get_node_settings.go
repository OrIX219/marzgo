package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type GetNodeSettings struct{}

func (GetNodeSettings) Body() RequestBody {
	return nil
}

func (GetNodeSettings) Query() QueryParams {
	return nil
}

func (GetNodeSettings) Headers() Headers {
	return nil
}

// GetNodeSettings returns node settings.
func (c *Client) GetNodeSettings() (models.NodeSettings, error) {
	var settings models.NodeSettings
	resp, err := c.Request("GET", "api/node/settings", GetNodeSettings{})
	if err != nil {
		return settings, err
	}

	err = json.Unmarshal(resp, &settings)

	return settings, err
}
