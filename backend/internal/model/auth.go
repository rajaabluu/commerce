package model

type contextKey string

const AuthContextKey contextKey = "auth"

type Auth struct {
	ID uint
}

type AuthResponse struct {
	ID      uint   `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	Role    uint   `json:"role,omitempty"`
	Contact string `json:"contact,omitempty"`
	Address string `json:"address,omitempty"`
}
