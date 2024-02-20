package models

// ProxyHosts represents hosts devided in inbounds.
// Can be cast to models.ModifyHosts if needed.
type ProxyHosts map[string][]ProxyHost

type ProxyHost struct {
	Remark        string               `json:"remark"`
	Address       string               `json:"address"`
	Port          *int                 `json:"port,omitempty"`
	SNI           string               `json:"sni,omitempty"`
	Host          string               `json:"host,omitempty"`
	Path          string               `json:"path,omitempty"`
	Security      ProxyHostSecurity    `json:"security,omitempty"`
	ALPN          ProxyHostALPN        `json:"alpn,omitempty"`
	Fingerprint   ProxyHostFingerprint `json:"fingerprint,omitempty"`
	AllowInsecure bool                 `json:"allowinsecure,omitempty"`
	IsDisabled    bool                 `json:"is_disabled,omitempty"`
}

type ProxyHostSecurity string

const (
	ProxyHostSecurityInboundDefault ProxyHostSecurity = "inbound_default"
	ProxyHostSecurityNone           ProxyHostSecurity = "none"
	ProxyHostSecurityTLS            ProxyHostSecurity = "tls"
)

type ProxyHostALPN string

type ProxyHostFingerprint string
