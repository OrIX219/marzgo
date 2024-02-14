package api

import (
	"encoding/json"
	"fmt"

	"github.com/OrIX219/marzgo/api/models"
)

type UserSubscriptionInfo struct {
	Token string
}

func (UserSubscriptionInfo) Method() string {
	return "GET"
}

func (p UserSubscriptionInfo) Endpoint() string {
	return fmt.Sprintf("sub/%s/info", p.Token)
}

func (p UserSubscriptionInfo) Params() (RequestParams, error) {
	return nil, nil
}

func (p UserSubscriptionInfo) Headers() (RequestParams, error) {
	return nil, nil
}

func (c *Client) UserSubscriptionInfo(params UserSubscriptionInfo) (models.UserResponse, error) {
	var sub models.UserResponse
	resp, err := c.Request(params)
	if err != nil {
		return sub, err
	}

	err = json.Unmarshal(resp, &sub)

	return sub, err
}
