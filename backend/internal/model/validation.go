package model

type ValidationErr struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}
