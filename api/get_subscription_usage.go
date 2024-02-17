package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/OrIX219/marzgo/api/models"
)

type GetSubscriptionUsage struct {
	Start time.Time
	End   time.Time
}

func (p GetSubscriptionUsage) Body() (RequestBody, error) {
	return nil, nil
}

func (p GetSubscriptionUsage) Query() (QueryParams, error) {
	params := make(QueryParams)
	params.AddTimeNonZero("start", p.Start)
	params.AddTimeNonZero("end", p.End)

	return params, nil
}

func (GetSubscriptionUsage) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) GetSubscriptionUsage(token string, params GetSubscriptionUsage) (models.UserUsagesResponse, error) {
	var usages models.UserUsagesResponse
	resp, err := c.Request("GET", fmt.Sprintf("sub/%s/usage", token), params)
	if err != nil {
		return usages, err
	}

	err = json.Unmarshal(resp, &usages)

	return usages, err
}
