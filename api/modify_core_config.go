package api

import (
	"encoding/json"

	"github.com/OrIX219/marzgo/api/models"
)

type ModifyCoreConfig models.CoreConfigModify

func (p ModifyCoreConfig) Body() RequestBody {
	return models.CoreConfigModify(p)
}

func (ModifyCoreConfig) Query() QueryParams {
	return nil
}

func (ModifyCoreConfig) Headers() Headers {
	return nil
}

// ModifyCoreConfig sets core config to provided one and returns it back.
func (c *Client) ModifyCoreConfig(config ModifyCoreConfig) (json.RawMessage, error) {
	return c.Request("PUT", "api/core/config", config)
}
