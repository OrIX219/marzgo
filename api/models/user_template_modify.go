package models

import "encoding/json"

type UserTemplateModify struct {
	Name           string       `json:"name,omitempty"`
	DataLimit      *int         `json:"data_limit,omitempty"`
	ExpireDuration *int         `json:"expire_duration,omitempty"`
	UsernamePrefix *string      `json:"username_prefix,omitempty"`
	UsernameSuffix *string      `json:"username_suffix,omitempty"`
	Inbounds       UserInbounds `json:"inbounds,omitempty"`
}

func (p UserTemplateModify) JSONEncoded() []byte {
	buf, _ := json.Marshal(p)
	return buf
}

func (p UserTemplateModify) URLEncoded() string {
	return ""
}
