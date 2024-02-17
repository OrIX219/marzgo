package models

import (
	"encoding/json"
)

type AdminModify struct {
	Password       string `json:"password,omitempty"`
	IsSudo         bool   `json:"is_sudo"`
	TelegramId     *int   `json:"telegram_id,omitempty"`
	DiscordWebhook string `json:"discord_webhook,omitempty"`
}

func (p AdminModify) JSONEncoded() []byte {
	buf, _ := json.Marshal(p)
	return buf
}

func (p AdminModify) URLEncoded() string {
	return ""
}
