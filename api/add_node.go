package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type AddNode models.NodeCreate

func (c AddNode) Body() (RequestBody, error) {
	return models.NodeCreate(c), nil
}

func (AddNode) Query() (QueryParams, error) {
	return nil, nil
}

func (c AddNode) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) AddNode(params AddNode) (models.NodeResponse, error) {
	var node models.NodeResponse
	resp, err := c.Request("POST", "api/node", params)
	if err != nil {
		return node, err
	}

	err = json.Unmarshal(resp, &node)

	return node, err
}
