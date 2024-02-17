package models

type Admin struct {
	Username       string `json:"username"`
	IsSudo         bool   `json:"is_sudo"`
	TelegramId     *int   `json:"telegram_id,omitempty"`
	DiscordWebhook string `json:"discord_webhook,omitempty"`
}
