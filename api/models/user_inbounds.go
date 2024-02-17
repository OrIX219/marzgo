package models

type UserInbounds struct {
	Vless       []VlessInbound       `json:"vless,omitempty"`
	Vmess       []VmessInbound       `json:"vmess,omitempty"`
	Trojan      []TrojanInbound      `json:"trojan,omitempty"`
	Shadowsocks []ShadowsocksInbound `json:"shadowsocks,omitempty"`
}
