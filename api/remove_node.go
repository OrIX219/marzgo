package api

import (
	"fmt"
)

type RemoveNode struct{}

func (c RemoveNode) Body() RequestBody {
	return nil
}

func (RemoveNode) Query() QueryParams {
	return nil
}

func (c RemoveNode) Headers() Headers {
	return nil
}

// RemoveNode removes a node with a specified id.
func (c *Client) RemoveNode(id int) error {
	_, err := c.Request("DELETE", fmt.Sprintf("api/node/%d", id), RemoveNode{})
	return err
}
