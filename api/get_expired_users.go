package api

import (
	"encoding/json"
	"time"
)

type GetExpiredUsers struct {
	ExpiredBefore time.Time
	ExpiredAfter  time.Time
}

func (GetExpiredUsers) Body() RequestBody {
	return nil
}

func (p GetExpiredUsers) Query() QueryParams {
	params := make(QueryParams)
	params.AddTimeNonZero("expired_before", p.ExpiredBefore)
	params.AddTimeNonZero("expired_after", p.ExpiredAfter)

	return params
}

func (GetExpiredUsers) Headers() Headers {
	return nil
}

// GetExpiredUsers returns usernames of all users
// whose subscriptions had expired in specified period of time
// (default all time).
func (c *Client) GetExpiredUsers(params GetExpiredUsers) ([]string, error) {
	users := []string{}
	resp, err := c.Request("GET", "api/users/expired", params)
	if err != nil {
		return users, err
	}

	err = json.Unmarshal(resp, &users)

	return users, err
}
