package models

type UserResponseInbounds struct {
	Vless       []VlessInbound       `json:"vless"`
	Vmess       []VmessInbound       `json:"vmess"`
	Trojan      []TrojanInbound      `json:"trojan"`
	Shadowsocks []ShadowsocksInbound `json:"shadowsocks"`
}

type VlessInbound string

type VmessInbound string

type TrojanInbound string

type ShadowsocksInbound string
