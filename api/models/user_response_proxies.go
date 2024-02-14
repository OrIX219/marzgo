package models

type UserResponseProxies struct {
	Vless       *VlessProxy       `json:"vless,omitempty"`
	Vmess       *VmessProxy       `json:"vmess,omitempty"`
	Trojan      *TrojanProxy      `json:"trojan,omitempty"`
	Shadowsocks *ShadowsocksProxy `json:"shadowsocks,omitempty"`
}

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
