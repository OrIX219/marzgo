package models

type UserTemplateResponse struct {
	Name           string       `json:"name,omitempty"`
	DataLimit      int          `json:"data_limit,omitempty"`
	ExpireDuration int          `json:"expire_duration,omitempty"`
	UsernamePrefix string       `json:"username_prefix"`
	UsernameSuffix string       `json:"username_suffix"`
	Inbounds       UserInbounds `json:"inbounds,omitempty"`
	Id             int          `json:"id"`
}
