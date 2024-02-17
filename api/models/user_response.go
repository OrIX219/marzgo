package models

type UserResponse struct {
	Proxies                UserProxies                `json:"proxies"`
	Expire                 int                        `json:"expire,omitempty"`
	DataLimit              int                        `json:"data_limit,omitempty"`
	DataLimitResetStrategy UserDataLimitResetStrategy `json:"data_limit_reset_strategy,omitempty"`
	Inbounds               UserInbounds               `json:"inbounds,omitempty"`
	Note                   string                     `json:"note,omitempty"`
	SubUpdatedAt           *Time                      `json:"sub_updated_at,omitempty"`
	SubLastUserAgent       string                     `json:"sub_last_user_agent,omitempty"`
	OnlineAt               *Time                      `json:"online_at,omitempty"`
	OnHoldExpireDuration   int                        `json:"on_hold_expire_duration,omitempty"`
	OnHoldTimeout          *Time                      `json:"on_hold_timeout,omitempty"`
	Username               string                     `json:"username"`
	Status                 UserStatus                 `json:"status"`
	UsedTraffic            int                        `json:"used_traffic"`
	LifetimeUsedTraffic    int                        `json:"lifetime_used_traffic,omitempty"`
	CreatedAt              *Time                      `json:"created_at"`
	Links                  []string                   `json:"links,omitempty"`
	SubscriptionURL        string                     `json:"subscription_url,omitempty"`
	ExcludedInbounds       UserInbounds               `json:"excluded_inbounds,omitempty"`
}
