package models

type UsersResponse struct {
	Total int            `json:"total"`
	Users []UserResponse `json:"users"`
}
