package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type GetSystemStats struct{}

func (GetSystemStats) Body() RequestBody {
	return nil
}

func (GetSystemStats) Query() QueryParams {
	return nil
}

func (GetSystemStats) Headers() Headers {
	return nil
}

// GetSystemStats returns system stats.
func (c *Client) GetSystemStats() (models.SystemStats, error) {
	var stats models.SystemStats
	resp, err := c.Request("GET", "api/system", GetSystemStats{})
	if err != nil {
		return stats, err
	}

	err = json.Unmarshal(resp, &stats)

	return stats, err
}
