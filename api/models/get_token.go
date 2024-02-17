package models

import "net/url"

type GetToken struct {
	GrantType    string
	Username     string
	Password     string
	Scope        string
	ClientId     string
	ClientSecret string
}

func (p GetToken) JSONEncoded() []byte {
	return []byte(nil)
}

func (p GetToken) URLEncoded() string {
	values := url.Values{}
	if p.GrantType != "" {
		values.Set("grant_type", p.GrantType)
	}
	values.Set("username", p.Username)
	values.Set("password", p.Password)
	if p.Scope != "" {
		values.Set("scope", p.Scope)
	}
	if p.ClientId != "" {
		values.Set("client_id", p.ClientId)
	}
	if p.ClientSecret != "" {
		values.Set("client_secret", p.ClientSecret)
	}

	return values.Encode()
}
