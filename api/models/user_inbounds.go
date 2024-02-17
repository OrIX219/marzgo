package models

type UserInbounds struct {
	Vless       []VlessInbound       `json:"vless"`
	Vmess       []VmessInbound       `json:"vmess"`
	Trojan      []TrojanInbound      `json:"trojan"`
	Shadowsocks []ShadowsocksInbound `json:"shadowsocks"`
}
