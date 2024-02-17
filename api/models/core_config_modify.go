package models

import "encoding/json"

type CoreConfigModify json.RawMessage

func (p CoreConfigModify) JSONEncoded() []byte {
	return p
}

func (p CoreConfigModify) URLEncoded() string {
	return ""
}
