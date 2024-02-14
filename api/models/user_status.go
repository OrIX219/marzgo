package models

type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusDisabled UserStatus = "disabled"
	UserStatusExpired  UserStatus = "expired"
	UserStatusLimited  UserStatus = "limited"
)
