package common

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Meta    any    `json:"meta"`
	Data    any    `json:"data"`
}
