package models

type NodeResponse struct {
	Name             string     `json:"name"`
	Address          string     `json:"address"`
	Port             int        `json:"port"`
	APIPort          int        `json:"api_port"`
	UsageCoefficient float64    `json:"usage_coefficient"`
	Id               int        `json:"id"`
	XrayVersion      string     `json:"xray_version"`
	Status           NodeStatus `json:"status"`
	Message          string     `json:"message,omitempty"`
}
