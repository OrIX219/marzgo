package models

import "github.com/google/uuid"

type VlessProxy struct {
	Id   uuid.UUID `json:"id"`
	Flow VlessFlow `json:"flow"`
}

type VlessFlow string

const (
	VlessFlowNone             VlessFlow = ""
	VlessFlowXTLS_RPRX_VISION VlessFlow = "xtls-rprx-vision"
)

type VmessProxy struct {
	Id uuid.UUID `json:"id"`
}

type TrojanProxy struct {
	Password string `json:"password"`
	Flow     string `json:"flow"`
}

type ShadowsocksProxy struct {
	Password string `json:"password"`
	Method   string `json:"method"`
}
