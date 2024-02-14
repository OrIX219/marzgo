package models

type UserUsagesResponse struct {
	Username string              `json:"username"`
	Usages   []UserUsageResponse `json:"usages"`
}

type UserUsageResponse struct {
	NodeId      *int   `json:"node_id"`
	NodeName    string `json:"node_name"`
	UsedTraffic uint64 `json:"used_traffic"`
}
