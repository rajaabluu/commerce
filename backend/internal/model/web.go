package model

import "github.com/golang-jwt/jwt/v5"

type AuthClaims struct {
	ID uint `json:"id,omitempty"`
	jwt.RegisteredClaims
}

type Response[T any] struct {
	Message string `json:"message,omitempty"`
	Data    T      `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
	Error   any    `json:"error,omitempty"`
}
