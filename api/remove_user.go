package api

import (
	"fmt"
)

type RemoveUser struct{}

func (RemoveUser) Body() (RequestBody, error) {
	return nil, nil
}

func (RemoveUser) Query() (QueryParams, error) {
	return nil, nil
}

func (RemoveUser) Headers() (Headers, error) {
	return nil, nil
}

// RemoveUser removes a user with a specified username.
func (c *Client) RemoveUser(username string) error {
	_, err := c.Request("DELETE", fmt.Sprintf("api/user/%s", username), RemoveUser{})
	return err
}
