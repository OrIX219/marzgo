package api

import (
	"encoding/json"
	"time"

	"github.com/OrIX219/marzgo/api/models"
)

type DeleteExpiredUsers struct {
	ExpiredBefore time.Time
	ExpiredAfter  time.Time
}

func (DeleteExpiredUsers) Body() (RequestBody, error) {
	return nil, nil
}

func (p DeleteExpiredUsers) Query() (QueryParams, error) {
	params := make(QueryParams)
	params.AddTimeNonZero("expired_before", p.ExpiredBefore)
	params.AddTimeNonZero("expired_after", p.ExpiredAfter)

	return params, nil
}

func (DeleteExpiredUsers) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) DeleteExpiredUsers(params DeleteExpiredUsers) error {
	_, err := c.Request("DELETE", "api/users/expired", params)
	return err
}
