package models

type UserCreateStatus string

const (
	UserCreateStatusActive UserCreateStatus = "active"
	UserCreateStatusOnHold UserCreateStatus = "on_hold"
)
