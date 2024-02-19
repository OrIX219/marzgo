package models

import "encoding/json"

type NodeCreate struct {
	Name             string   `json:"name"`
	Address          string   `json:"address"`
	Port             *int     `json:"port,omitempty"`
	APIPort          *int     `json:"api_port,omitempty"`
	UsageCoefficient *float64 `json:"usage_coefficient,omitempty"`
	AddAsNewHost     *bool    `json:"add_as_new_host,omitempty"`
}

func (p NodeCreate) JSONEncoded() []byte {
	buf, _ := json.Marshal(p)
	return buf
}

func (p NodeCreate) URLEncoded() string {
	return ""
}
