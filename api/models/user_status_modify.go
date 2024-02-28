package models

type UserStatusModify string

const (
	UserStatusModifyActive   UserStatusModify = "active"
	UserStatusModifyOnHold   UserStatusModify = "on_hold"
	UserStatusModifyDisabled UserStatusModify = "disabled"
)
