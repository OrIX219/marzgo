package models

import "encoding/json"

type NodeModify struct {
	Name             *string     `json:"name,omitempty"`
	Address          *string     `json:"address,omitempty"`
	Port             *int        `json:"port,omitempty"`
	APIPort          *int        `json:"api_port,omitempty"`
	UsageCoefficient *float64    `json:"usage_coefficient,omitempty"`
	Status           *NodeStatus `json:"status,omitempty"`
}

func (p NodeModify) JSONEncoded() []byte {
	buf, _ := json.Marshal(p)
	return buf
}

func (p NodeModify) URLEncoded() string {
	return ""
}
