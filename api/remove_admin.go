package api

import (
	"fmt"
)

type RemoveAdminParams struct {
	Username string
}

func (RemoveAdminParams) Method() string {
	return "DELETE"
}

func (p RemoveAdminParams) Endpoint() string {
	return fmt.Sprintf("api/admin/%s", p.Username)
}

func (c RemoveAdminParams) Params() (RequestParams, error) {
	return nil, nil
}

func (c *Client) RemoveAdmin(params RemoveAdminParams) error {
	_, err := c.Request(params)
	if err != nil {
		return err
	}

	return err
}
