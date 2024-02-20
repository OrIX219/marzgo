package api

import (
	"fmt"
)

type RemoveUserTemplate struct{}

func (c RemoveUserTemplate) Body() RequestBody {
	return nil
}

func (RemoveUserTemplate) Query() QueryParams {
	return nil
}

func (c RemoveUserTemplate) Headers() Headers {
	return nil
}

// RemoveUserTemplate removes a user template with a specified id.
func (c *Client) RemoveUserTemplate(id int) error {
	_, err := c.Request("DELETE", fmt.Sprintf("api/user_template/%d", id), RemoveUserTemplate{})
	return err
}
