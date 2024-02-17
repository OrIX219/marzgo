package models

import (
	"encoding/json"
)

type AdminCreate struct {
	Username       string `json:"username,omitempty"`
	Password       string `json:"password,omitempty"`
	IsSudo         bool   `json:"is_sudo"`
	TelegramId     *int   `json:"telegram_id,omitempty"`
	DiscordWebhook string `json:"discord_webhook,omitempty"`
}

func (p AdminCreate) JSONEncoded() []byte {
	buf, _ := json.Marshal(p)
	return buf
}

func (p AdminCreate) URLEncoded() string {
	return ""
}
