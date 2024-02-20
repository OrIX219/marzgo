package api

type RestartCore struct{}

func (RestartCore) Body() RequestBody {
	return nil
}

func (RestartCore) Query() QueryParams {
	return nil
}

func (RestartCore) Headers() Headers {
	return nil
}

// RestartCore restarts core.
func (c *Client) RestartCore() error {
	_, err := c.Request("POST", "api/core/restart", RestartCore{})
	return err
}
