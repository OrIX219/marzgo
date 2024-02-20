package api

type ResetUsersDataUsage struct{}

func (c ResetUsersDataUsage) Body() RequestBody {
	return nil
}

func (ResetUsersDataUsage) Query() QueryParams {
	return nil
}

func (c ResetUsersDataUsage) Headers() Headers {
	return nil
}

// ResetUsersDataUsage resets data usage of all users.
func (c *Client) ResetUsersDataUsage() error {
	_, err := c.Request("POST", "api/users/reset", ResetUsersDataUsage{})
	return err
}
