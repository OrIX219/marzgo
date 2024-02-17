package models

type UserResponse struct {
	Proxies                UserProxies                `json:"proxies"`
	Expire                 uint64                     `json:"expire,omitempty"`
	DataLimit              uint64                     `json:"data_limit,omitempty"`
	DataLimitResetStrategy UserDataLimitResetStrategy `json:"data_limit_reset_strategy,omitempty"`
	Inbounds               UserInbounds               `json:"inbounds,omitempty"`
	Note                   string                     `json:"note,omitempty"`
	SubUpdatedAt           Time                       `json:"sub_updated_at,omitempty"`
	SubLastUserAgent       string                     `json:"sub_last_user_agent,omitempty"`
	OnlineAt               Time                       `json:"online_at,omitempty"`
	OnHoldExpireDuration   uint64                     `json:"on_hold_expire_duration,omitempty"`
	OnHoldTimeout          string                     `json:"on_hold_timeout,omitempty"`
	Username               string                     `json:"username"`
	Status                 UserStatus                 `json:"status"`
	UsedTraffic            uint64                     `json:"used_traffic"`
	LifetimeUsedTraffic    uint64                     `json:"lifetime_used_traffic,omitempty"`
	CreatedAt              Time                       `json:"created_at"`
	Links                  []string                   `json:"links,omitempty"`
	SubscriptionURL        string                     `json:"subscription_url,omitempty"`
	ExcludedInbounds       UserInbounds               `json:"excluded_inbounds,omitempty"`
}
