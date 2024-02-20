package api

import (
	"fmt"
)

type ReconnectNode struct{}

func (c ReconnectNode) Body() (RequestBody, error) {
	return nil, nil
}

func (ReconnectNode) Query() (QueryParams, error) {
	return nil, nil
}

func (c ReconnectNode) Headers() (Headers, error) {
	return nil, nil
}

// ReconnectNode reconnects main server to a node with a specified id.
func (c *Client) ReconnectNode(id int) error {
	_, err := c.Request("POST", fmt.Sprintf("api/node/%d/reconnect", id), ReconnectNode{})
	return err
}
