package models

type NodeSettings struct {
	MinNodeVersion string `json:"min_node_version,omitempty"`
	Certificate    string `json:"certificate"`
}
