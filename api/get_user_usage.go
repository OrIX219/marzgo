package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/OrIX219/marzgo/api/models"
)

type GetUserUsage struct {
	Token string
	Start time.Time
	End   time.Time
}

func (GetUserUsage) Method() string {
	return "GET"
}

func (p GetUserUsage) Endpoint() string {
	return fmt.Sprintf("sub/%s/usage", p.Token)
}

func (p GetUserUsage) Params() (RequestParams, error) {
	params := make(RequestParams)
	params.AddTimeNonZero("start", p.Start)
	params.AddTimeNonZero("end", p.End)

	return params, nil
}

func (GetUserUsage) Headers() (RequestParams, error) {
	return nil, nil
}

func (c *Client) GetUserUsage(params GetUserUsage) (models.UserUsagesResponse, error) {
	var usages models.UserUsagesResponse
	resp, err := c.Request(params)
	if err != nil {
		return usages, err
	}

	err = json.Unmarshal(resp, &usages)

	return usages, err
}
