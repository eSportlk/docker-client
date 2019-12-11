package models

type Response struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code, omitempty"`
}

const (
	PARAMETER_NOT_FOUND int = 400
	SERVER_ERROR        int = 500
	SUCCESS             int = 200
)
