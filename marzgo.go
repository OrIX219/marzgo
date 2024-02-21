// Package marzgo provides a way to interact with Marzban's HTTP API
package marzgo

import (
	"errors"
	"net/http"

	"github.com/OrIX219/marzgo/api"
)

// MarzbanAPI allows you to interact with Marzban API
type MarzbanAPI struct {
	api.Client
}

// NewMarzbanAPI creates new MarzbanAPI instance.
// apiBaseURL is part of url prior to 'api/...'.
// Requires username and password to request admin token.
func NewMarzbanAPI(apiBaseURL, username, password string) (*MarzbanAPI, error) {
	return NewMarzbanAPIWithHTTP(apiBaseURL, username, password, &http.Client{})
}

// NewMarzbanAPIWithHTTP creates new MarzbanAPI instance.
// apiBaseURL is part of url prior to 'api/...'.
// Requires username and password to request admin token.
// Additionally accepts custom HTTPClient
func NewMarzbanAPIWithHTTP(apiBaseURL, username, password string, httpClient api.HTTPClient) (*MarzbanAPI, error) {
	if len(apiBaseURL) == 0 {
		return nil, errors.New("Invalid apiBaseURL")
	}
	if apiBaseURL[len(apiBaseURL)-1] != '/' {
		apiBaseURL += "/"
	}
	marzban := &MarzbanAPI{
		Client: api.NewClient(apiBaseURL, username, password, httpClient),
	}

	token, err := marzban.AdminToken()
	if err != nil {
		return nil, err
	}
	marzban.SetToken(token.AccessToken)

	return marzban, nil
}
