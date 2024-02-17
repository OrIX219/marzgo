package api

import (
	"fmt"
)

type RemoveAdmin struct {
	Username string
}

func (RemoveAdmin) Method() string {
	return "DELETE"
}

func (p RemoveAdmin) Endpoint() string {
	return fmt.Sprintf("api/admin/%s", p.Username)
}

func (RemoveAdmin) Body() (RequestBody, error) {
	return nil, nil
}

func (RemoveAdmin) Query() (QueryParams, error) {
	return nil, nil
}

func (RemoveAdmin) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) RemoveAdmin(params RemoveAdmin) error {
	_, err := c.Request(params)
	if err != nil {
		return err
	}

	return err
}
