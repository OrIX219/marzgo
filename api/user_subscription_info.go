package api

import (
	"encoding/json"
	"fmt"

	"github.com/OrIX219/marzgo/api/models"
)

type UserSubscriptionInfo struct{}

func (UserSubscriptionInfo) Body() RequestBody {
	return nil
}

func (UserSubscriptionInfo) Query() QueryParams {
	return nil
}

func (UserSubscriptionInfo) Headers() Headers {
	return nil
}

// UserSubscriptionInfo returns a user corresponding to a specified sub token.
func (c *Client) UserSubscriptionInfo(token string) (models.UserResponse, error) {
	var sub models.UserResponse
	resp, err := c.Request("GET", fmt.Sprintf("sub/%s/info", token), UserSubscriptionInfo{})
	if err != nil {
		return sub, err
	}

	err = json.Unmarshal(resp, &sub)

	return sub, err
}
