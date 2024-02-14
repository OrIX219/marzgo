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

func (RemoveAdmin) Params() (RequestParams, error) {
	return nil, nil
}

func (RemoveAdmin) Headers() (RequestParams, error) {
	return nil, nil
}

func (c *Client) RemoveAdmin(params RemoveAdmin) error {
	_, err := c.Request(params)
	if err != nil {
		return err
	}

	return err
}
