package api

import (
	"encoding/json"
	"time"

	"github.com/OrIX219/marzgo/api/models"
)

type GetExpiredUsers struct {
	ExpiredBefore time.Time
	ExpiredAfter  time.Time
}

func (GetExpiredUsers) Body() (RequestBody, error) {
	return nil, nil
}

func (p GetExpiredUsers) Query() (QueryParams, error) {
	params := make(QueryParams)
	params.AddTimeNonZero("expired_before", p.ExpiredBefore)
	params.AddTimeNonZero("expired_after", p.ExpiredAfter)

	return params, nil
}

func (GetExpiredUsers) Headers() (Headers, error) {
	return nil, nil
}

func (c *Client) GetExpiredUsers(params GetExpiredUsers) (models.UsersResponse, error) {
	var users models.UsersResponse
	resp, err := c.Request("GET", "api/users/expired", params)
	if err != nil {
		return users, err
	}

	err = json.Unmarshal(resp, &users)

	return users, err
}
