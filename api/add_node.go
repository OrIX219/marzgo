package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type AddNode models.NodeCreate

func (c AddNode) Body() RequestBody {
	return models.NodeCreate(c)
}

func (AddNode) Query() QueryParams {
	return nil
}

func (c AddNode) Headers() Headers {
	return nil
}

// AddNode adds a new Marzban node with specified params and
// returns resulting node.
//
// Required params are Name and Address.
func (c *Client) AddNode(params AddNode) (models.NodeResponse, error) {
	var node models.NodeResponse
	resp, err := c.Request("POST", "api/node", params)
	if err != nil {
		return node, err
	}

	err = json.Unmarshal(resp, &node)

	return node, err
}
