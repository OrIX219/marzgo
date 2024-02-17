package models

type ProxyInbounds struct {
	Vmess       []ProxyInbound `json:"vmess,omitempty"`
	Vless       []ProxyInbound `json:"vless,omitempty"`
	Trojan      []ProxyInbound `json:"trojan,omitempty"`
	Shadowsocks []ProxyInbound `json:"shadowsocks,omitempty"`
}

type ProxyInbound struct {
	Tag      string    `json:"tag"`
	Protocol ProxyType `json:"protocol"`
	Network  string    `json:"network"`
	TLS      string    `json:"tls"`
	Port     int       `json:"port"`
}

type ProxyType string

const (
	ProxyTypeVmess       ProxyType = "vmess"
	ProxyTypeVless       ProxyType = "vless"
	ProxyTypeTrojan      ProxyType = "trojan"
	ProxyTypeShadowsocks ProxyType = "shadowsocks"
)
