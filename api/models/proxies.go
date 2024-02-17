package models

type VlessProxy struct {
	Id   string `json:"id"`
	Flow string `json:"flow"`
}

type VmessProxy struct {
	Id string `json:"id"`
}

type TrojanProxy struct {
	Password string `json:"password"`
	Flow     string `json:"flow"`
}

type ShadowsocksProxy struct {
	Password string `json:"password"`
	Method   string `json:"method"`
}
