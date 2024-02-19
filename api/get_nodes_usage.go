package api

import (
	"encoding/json"
	"time"

	"github.com/OrIX219/marzgo/api/models"
)

type GetNodesUsage struct {
	Start time.Time
	End   time.Time
}

func (p GetNodesUsage) Body() (RequestBody, error) {
	return nil, nil
}

func (p GetNodesUsage) Query() (QueryParams, error) {
	params := make(QueryParams)
	params.AddTimeNonZero("start", p.Start)
	params.AddTimeNonZero("end", p.End)

	return params, nil
}

func (GetNodesUsage) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) GetNodesUsage(params GetNodesUsage) (models.NodesUsageResponse, error) {
	var usages models.NodesUsageResponse
	resp, err := c.Request("GET", "api/nodes/usage", params)
	if err != nil {
		return usages, err
	}

	err = json.Unmarshal(resp, &usages)

	return usages, err
}
