package api

import (
	"fmt"
)

type RemoveAdmin struct{}

func (RemoveAdmin) Body() (RequestBody, error) {
	return nil, nil
}

func (RemoveAdmin) Query() (QueryParams, error) {
	return nil, nil
}

func (RemoveAdmin) Headers() (Headers, error) {
	return nil, nil
}

// RemoveAdmin removes an admin with a specified username.
func (c *Client) RemoveAdmin(username string) error {
	_, err := c.Request("DELETE", fmt.Sprintf("api/admin/%s", username), RemoveAdmin{})
	if err != nil {
		return err
	}

	return err
}
