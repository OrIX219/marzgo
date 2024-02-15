package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/OrIX219/marzgo/api/models"
)

type GetSubscriptionUsage struct {
	Token string
	Start time.Time
	End   time.Time
}

func (GetSubscriptionUsage) Method() string {
	return "GET"
}

func (p GetSubscriptionUsage) Endpoint() string {
	return fmt.Sprintf("sub/%s/usage", p.Token)
}

func (p GetSubscriptionUsage) Body() (BodyParams, error) {
	return nil, nil
}

func (p GetSubscriptionUsage) Url() (UrlParams, error) {
	params := make(UrlParams)
	params.AddTimeNonZero("start", p.Start)
	params.AddTimeNonZero("end", p.End)

	return params, nil
}

func (GetSubscriptionUsage) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) GetSubscriptionUsage(params GetSubscriptionUsage) (models.UserUsagesResponse, error) {
	var usages models.UserUsagesResponse
	resp, err := c.Request(params)
	if err != nil {
		return usages, err
	}

	err = json.Unmarshal(resp, &usages)

	return usages, err
}
