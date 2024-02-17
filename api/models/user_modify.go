package models

import "encoding/json"

type UserModify struct {
	Proxies                UserProxies                `json:"proxies,omitempty"`
	Expire                 uint64                     `json:"expire,omitempty"`
	DataLimit              uint64                     `json:"data_limit,omitempty"`
	DataLimitResetStrategy UserDataLimitResetStrategy `json:"data_limit_reset_strategy,omitempty"`
	Inbounds               UserInbounds               `json:"inbounds,omitempty"`
	Note                   string                     `json:"note,omitempty"`
	SubUpdatedAt           *Time                      `json:"sub_updated_at,omitempty"`
	SubLastUserAgent       string                     `json:"sub_last_user_agent,omitempty"`
	OnlineAt               *Time                      `json:"online_at,omitempty"`
	OnHoldExpireDuration   uint64                     `json:"on_hold_expire_duration,omitempty"`
	OnHoldTimeout          *Time                      `json:"on_hold_timeout,omitempty"`
	Status                 UserStatusModify           `json:"status,omitempty"`
}

func (p UserModify) JSONEncoded() []byte {
	buf, _ := json.Marshal(p)
	return buf
}

func (p UserModify) URLEncoded() string {
	return ""
}
