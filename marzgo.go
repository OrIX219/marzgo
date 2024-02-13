package marzgo

import (
	"net/http"

	"github.com/OrIX219/marzgo/api"
)

type MarzbanAPI struct {
	api.Client
}

func NewMarzbanAPI(apiBaseURL, username, password string) (*MarzbanAPI, error) {
	return NewMarzbanAPIWithHTTP(apiBaseURL, username, password, &http.Client{})
}

func NewMarzbanAPIWithHTTP(apiBaseURL, username, password string, httpClient api.HTTPClient) (*MarzbanAPI, error) {
	marzban := &MarzbanAPI{
		Client: api.Client{
			APIBaseUrl: apiBaseURL,
			Username:   username,
			Password:   password,
			HTTPClient: httpClient,
		},
	}

	token, err := marzban.AdminToken()
	if err != nil {
		return nil, err
	}

	marzban.Client.Token = token.AccessToken

	return marzban, nil
}
