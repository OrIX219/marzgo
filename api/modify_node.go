package api

import (
	"encoding/json"
	"fmt"

	"github.com/OrIX219/marzgo/api/models"
)

type ModifyNode models.NodeModify

func (c ModifyNode) Body() (RequestBody, error) {
	return models.NodeModify(c), nil
}

func (ModifyNode) Query() (QueryParams, error) {
	return nil, nil
}

func (c ModifyNode) Headers() (Headers, error) {
	return nil, nil
}

// ModifyNode modifies a node with a specified id and returns a resulting node.
func (c *Client) ModifyNode(id int, params ModifyNode) (models.NodeResponse, error) {
	var node models.NodeResponse
	resp, err := c.Request("PUT", fmt.Sprintf("api/node/%d", id), params)
	if err != nil {
		return node, err
	}

	err = json.Unmarshal(resp, &node)

	return node, err
}
