package api

import (
	"fmt"
)

type RemoveAdmin struct{}

func (RemoveAdmin) Body() RequestBody {
	return nil
}

func (RemoveAdmin) Query() QueryParams {
	return nil
}

func (RemoveAdmin) Headers() Headers {
	return nil
}

// RemoveAdmin removes an admin with a specified username.
func (c *Client) RemoveAdmin(username string) error {
	_, err := c.Request("DELETE", fmt.Sprintf("api/admin/%s", username), RemoveAdmin{})
	if err != nil {
		return err
	}

	return err
}
