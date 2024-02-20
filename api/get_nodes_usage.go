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

func (p GetNodesUsage) Body() RequestBody {
	return nil
}

func (p GetNodesUsage) Query() QueryParams {
	params := make(QueryParams)
	params.AddTimeNonZero("start", p.Start)
	params.AddTimeNonZero("end", p.End)

	return params
}

func (GetNodesUsage) Headers() Headers {
	return nil
}

// GetNodesUsage returns usages of nodes for a specified period of time (default all time).
func (c *Client) GetNodesUsage(params GetNodesUsage) (models.NodesUsageResponse, error) {
	var usages models.NodesUsageResponse
	resp, err := c.Request("GET", "api/nodes/usage", params)
	if err != nil {
		return usages, err
	}

	err = json.Unmarshal(resp, &usages)

	return usages, err
}
