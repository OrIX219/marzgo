package models

type UserDataLimitResetStrategy string

const (
	UserDataLimitResetStrategyNoReset UserDataLimitResetStrategy = "no_reset"
	UserDataLimitResetStrategyDay     UserDataLimitResetStrategy = "day"
	UserDataLimitResetStrategyWeek    UserDataLimitResetStrategy = "week"
	UserDataLimitResetStrategyMonth   UserDataLimitResetStrategy = "month"
	UserDataLimitResetStrategyYear    UserDataLimitResetStrategy = "year"
)
