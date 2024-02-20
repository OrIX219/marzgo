package api

import (
	"fmt"
)

type RemoveUser struct{}

func (RemoveUser) Body() RequestBody {
	return nil
}

func (RemoveUser) Query() QueryParams {
	return nil
}

func (RemoveUser) Headers() Headers {
	return nil
}

// RemoveUser removes a user with a specified username.
func (c *Client) RemoveUser(username string) error {
	_, err := c.Request("DELETE", fmt.Sprintf("api/user/%s", username), RemoveUser{})
	return err
}
