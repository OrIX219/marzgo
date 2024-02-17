package api

type ResetUsersDataUsage struct{}

func (c ResetUsersDataUsage) Body() (RequestBody, error) {
	return nil, nil
}

func (ResetUsersDataUsage) Query() (QueryParams, error) {
	return nil, nil
}

func (c ResetUsersDataUsage) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) ResetUsersDataUsage() error {
	_, err := c.Request("POST", "api/users/reset", ResetUsersDataUsage{})
	return err
}
