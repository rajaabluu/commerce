package middleware

import "github.com/rajaabluu/ershop-api/internal/service"

type Middleware struct {
	UserService *service.UserService
}

func NewMiddleware(userService *service.UserService) *Middleware {
	return &Middleware{
		UserService: userService,
	}
}
