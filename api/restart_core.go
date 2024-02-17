package api

type RestartCore struct{}

func (RestartCore) Body() (RequestBody, error) {
	return nil, nil
}

func (RestartCore) Query() (QueryParams, error) {
	return nil, nil
}

func (RestartCore) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) RestartCore() error {
	_, err := c.Request("POST", "api/core/restart", RestartCore{})
	return err
}
