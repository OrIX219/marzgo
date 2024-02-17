package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type ModifyCoreConfig models.CoreConfigModify

func (p ModifyCoreConfig) Body() (RequestBody, error) {
	return models.CoreConfigModify(p), nil
}

func (ModifyCoreConfig) Query() (QueryParams, error) {
	return nil, nil
}

func (ModifyCoreConfig) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) ModifyCoreConfig(config ModifyCoreConfig) (json.RawMessage, error) {
	return c.Request("PUT", "api/core/config", config)
}
