package models

type VlessProxy struct {
	Id   string    `json:"id,omitempty"`
	Flow VlessFlow `json:"flow,omitempty"`
}

type VlessFlow string

const (
	VlessFlowNone             VlessFlow = ""
	VlessFlowXTLS_RPRX_VISION VlessFlow = "xtls-rprx-vision"
)

type VmessProxy struct {
	Id string `json:"id,omitempty"`
}

type TrojanProxy struct {
	Password string `json:"password,omitempty"`
	Flow     string `json:"flow,omitempty"`
}

type ShadowsocksProxy struct {
	Password string `json:"password,omitempty"`
	Method   string `json:"method,omitempty"`
}
