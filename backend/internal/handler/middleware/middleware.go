package middleware

import "github.com/rajaabluu/ershop-api/internal/service"

type Middleware struct {
	CustomerService *service.CustomerService
}

func NewMiddleware(customerService *service.CustomerService) *Middleware {
	return &Middleware{
		CustomerService: customerService,
	}
}
