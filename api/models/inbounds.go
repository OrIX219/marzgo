package models

type VlessInbound string

const (
	VlessInboundTCPReality  = "VLESS TCP REALITY"
	VlessInboundGRPCReality = "VLESS GRPC REALITY"
)

type VmessInbound string

const (
	VmessInboundTCP = "VMESS TCP"
	VmessInboundWS  = "VMESS WEBSOCKET"
)

type TrojanInbound string

const (
	TrojanInboundWS = "TROJAN WEBSOCKET"
)

type ShadowsocksInbound string

const (
	ShadowsocksInboundTCP = "SHADOWSOCKS TCP"
)
