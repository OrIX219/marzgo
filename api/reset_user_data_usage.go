package api

import (
	"encoding/json"
	"fmt"

	"github.com/OrIX219/marzgo/api/models"
)

type ResetUserDataUsage struct{}

func (c ResetUserDataUsage) Body() RequestBody {
	return nil
}

func (ResetUserDataUsage) Query() QueryParams {
	return nil
}

func (c ResetUserDataUsage) Headers() Headers {
	return nil
}

// ResetUserDataUsage resets data usage of a user with a specified username and returns that user.
func (c *Client) ResetUserDataUsage(username string) (models.UserResponse, error) {
	var user models.UserResponse
	resp, err := c.Request("POST", fmt.Sprintf("api/user/%s/reset", username), ResetUserDataUsage{})
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(resp, &user)

	return user, err
}
