package api

import (
	"encoding/json"
	"fmt"

	"github.com/OrIX219/marzgo/api/models"
)

type GetNode struct{}

func (GetNode) Body() (RequestBody, error) {
	return nil, nil
}

func (GetNode) Query() (QueryParams, error) {
	return nil, nil
}

func (GetNode) Headers() (Headers, error) {
	return nil, nil
}

// GetNode returns a node with specified id.
func (c *Client) GetNode(id int) (models.NodeResponse, error) {
	var node models.NodeResponse
	resp, err := c.Request("GET", fmt.Sprintf("api/node/%d", id), GetNode{})
	if err != nil {
		return node, err
	}

	err = json.Unmarshal(resp, &node)

	return node, err
}
