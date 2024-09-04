package model

type contextKey string

const AuthContextKey contextKey = "auth"

type Auth struct {
	ID uint
}
