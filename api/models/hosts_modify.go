package models

import "encoding/json"

type HostsModify map[string][]ProxyHost

func (p HostsModify) JSONEncoded() []byte {
	buf, _ := json.Marshal(p)
	return buf
}

func (p HostsModify) URLEncoded() string {
	return ""
}
