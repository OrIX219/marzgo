package models

type NodesUsageResponse struct {
	Usages []NodeUsageResponse `json:"usages"`
}

type NodeUsageResponse struct {
	NodeId   int    `json:"node_id"`
	NodeName string `json:"node_name"`
	Uplink   int    `json:"uplink"`
	Downlink int    `json:"downlink"`
}
