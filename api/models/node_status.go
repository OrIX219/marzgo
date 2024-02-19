package models

type NodeStatus string

const (
	NodeStatusConnected  NodeStatus = "connected"
	NodeStatusConnecting NodeStatus = "connecting"
	NodeStatusError      NodeStatus = "error"
	NodeStatusDisabled   NodeStatus = "disabled"
)
