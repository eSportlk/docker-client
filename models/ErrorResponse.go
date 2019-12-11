package models

type ErrorResponse struct{
	Message string `json:"message,omitempty"`
	Code int `json:"code, omitempty"`
}
