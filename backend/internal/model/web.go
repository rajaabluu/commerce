package model

type Response[T any] struct {
	Message string `json:"message,omitempty"`
	Data    T      `json:"data"`
	Errors  any    `json:"errors,omitempty"`
}

type ResponseWithToken[T any] struct {
	*Response[T]
	Token string `json:"token,omitempty"`
}

type ErrResponse struct {
	Message string `json:"message,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}
