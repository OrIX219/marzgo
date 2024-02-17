package models

type SystemStats struct {
	Version                string  `json:"version"`
	MemTotal               int     `json:"mem_total"`
	MemUsed                int     `json:"mem_used"`
	CPUCores               int     `json:"cpu_cores"`
	CPUUsage               float64 `json:"cpu_usage"`
	TotalUser              int     `json:"total_user"`
	UsersActive            int     `json:"users_active"`
	IncomingBandwidth      int     `json:"incoming_bandwidth"`
	OutgoindBandwidth      int     `json:"outgoing_bandwidth"`
	IncomingBandwidthSpeed int     `json:"incoming_bandwidth_speed"`
	OutgoingBandwidthSpeed int     `json:"outgoing_bandwidth_speed"`
}
