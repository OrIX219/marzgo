package models

type Admin struct {
	Username string `json:"username"`
	IsSudo   bool   `json:"is_sudo"`
}
