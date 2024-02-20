package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type GetCoreStats struct{}

func (GetCoreStats) Body() RequestBody {
	return nil
}

func (GetCoreStats) Query() QueryParams {
	return nil
}

func (GetCoreStats) Headers() Headers {
	return nil
}

// GetCoreStats returns core stats.
func (c *Client) GetCoreStats() (models.CoreStats, error) {
	var stats models.CoreStats
	resp, err := c.Request("GET", "api/core", GetCoreStats{})
	if err != nil {
		return stats, err
	}

	err = json.Unmarshal(resp, &stats)

	return stats, err
}
