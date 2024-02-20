package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/OrIX219/marzgo/api/models"
)

type GetUserUsage struct {
	Start time.Time
	End   time.Time
}

func (p GetUserUsage) Body() (RequestBody, error) {
	return nil, nil
}

func (p GetUserUsage) Query() (QueryParams, error) {
	params := make(QueryParams)
	params.AddTimeNonZero("start", p.Start)
	params.AddTimeNonZero("end", p.End)

	return params, nil
}

func (GetUserUsage) Headers() (Headers, error) {
	return nil, nil
}

// GetUserUsage returns user's data usage for each node
// for a specified period of time (default all time).
func (c *Client) GetUserUsage(username string, params GetUserUsage) (models.UserUsagesResponse, error) {
	var usages models.UserUsagesResponse
	resp, err := c.Request("GET", fmt.Sprintf("api/user/%s/usage", username), params)
	if err != nil {
		return usages, err
	}

	err = json.Unmarshal(resp, &usages)

	return usages, err
}
