package models

type UserProxies struct {
	Vless       *VlessProxy       `json:"vless,omitempty"`
	Vmess       *VmessProxy       `json:"vmess,omitempty"`
	Trojan      *TrojanProxy      `json:"trojan,omitempty"`
	Shadowsocks *ShadowsocksProxy `json:"shadowsocks,omitempty"`
}
