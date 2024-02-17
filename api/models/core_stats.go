package models

type CoreStats struct {
	Version       string `json:"version"`
	Started       bool   `json:"started"`
	LogsWebsocket string `json:"logs_websocket"`
}
