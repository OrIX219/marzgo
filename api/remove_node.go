package api

import (
	"fmt"
)

type RemoveNode struct{}

func (c RemoveNode) Body() (RequestBody, error) {
	return nil, nil
}

func (RemoveNode) Query() (QueryParams, error) {
	return nil, nil
}

func (c RemoveNode) Headers() (Headers, error) {
	return nil, nil
}

// RemoveNode removes a node with a specified id.
func (c *Client) RemoveNode(id int) error {
	_, err := c.Request("DELETE", fmt.Sprintf("api/node/%d", id), RemoveNode{})
	return err
}
