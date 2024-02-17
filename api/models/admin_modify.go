package models

import (
	"encoding/json"
)

type AdminModify struct {
	Password string `json:"password,omitempty"`
	IsSudo   bool   `json:"is_sudo"`
}

func (p AdminModify) JSONEncoded() []byte {
	buf, _ := json.Marshal(p)
	return buf
}

func (p AdminModify) URLEncoded() string {
	return ""
}
