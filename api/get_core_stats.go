package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type GetCoreStats struct{}

func (GetCoreStats) Body() (RequestBody, error) {
	return nil, nil
}

func (GetCoreStats) Query() (QueryParams, error) {
	return nil, nil
}

func (GetCoreStats) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) GetCoreStats() (models.CoreStats, error) {
	var stats models.CoreStats
	resp, err := c.Request("GET", "api/core", GetCoreStats{})
	if err != nil {
		return stats, err
	}

	err = json.Unmarshal(resp, &stats)

	return stats, err
}
