package api

import (
	"encoding/json"
	"time"
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

// DeleteExpiredUsers deletes all users with expired subscription
// and returns their usernames.
func (c *Client) DeleteExpiredUsers(params DeleteExpiredUsers) ([]string, error) {
	users := []string{}
	resp, err := c.Request("DELETE", "api/users/expired", params)
	if err != nil {
		return users, err
	}

	err = json.Unmarshal(resp, &users)

	return users, err
}
