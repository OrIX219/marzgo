package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type GetNodes struct{}

func (GetNodes) Body() RequestBody {
	return nil
}

func (GetNodes) Query() QueryParams {
	return nil
}

func (GetNodes) Headers() Headers {
	return nil
}

// GetNodes returns all nodes.
func (c *Client) GetNodes() ([]models.NodeResponse, error) {
	nodes := []models.NodeResponse{}
	resp, err := c.Request("GET", "api/nodes", GetNodes{})
	if err != nil {
		return nodes, err
	}

	err = json.Unmarshal(resp, &nodes)

	return nodes, err
}
