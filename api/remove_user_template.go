package api

import (
	"fmt"
)

type RemoveUserTemplate struct{}

func (c RemoveUserTemplate) Body() (RequestBody, error) {
	return nil, nil
}

func (RemoveUserTemplate) Query() (QueryParams, error) {
	return nil, nil
}

func (c RemoveUserTemplate) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) RemoveUserTemplate(id int) error {
	_, err := c.Request("DELETE", fmt.Sprintf("api/user_template/%d", id), RemoveUserTemplate{})
	return err
}
