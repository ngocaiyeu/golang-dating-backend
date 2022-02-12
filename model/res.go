package model

type Response struct {
	StatusCode int         `json:"status-code,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
