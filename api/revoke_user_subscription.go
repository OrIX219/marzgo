package api

import (
	"encoding/json"
	"fmt"

	"github.com/OrIX219/marzgo/api/models"
)

type RevokeUserSubscription struct{}

func (c RevokeUserSubscription) Body() (RequestBody, error) {
	return nil, nil
}

func (RevokeUserSubscription) Query() (QueryParams, error) {
	return nil, nil
}

func (c RevokeUserSubscription) Headers() (Headers, error) {
	return nil, nil
}

// RevokeUserSubscription revokes subscription of a user with a specified username
// and returns a resulting user with a new subscription.
func (c *Client) RevokeUserSubscription(username string) (models.UserResponse, error) {
	var user models.UserResponse
	resp, err := c.Request("POST", fmt.Sprintf("api/user/%s/revoke_sub", username), RevokeUserSubscription{})
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(resp, &user)

	return user, err
}
