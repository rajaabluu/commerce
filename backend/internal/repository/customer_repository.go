package repository

import "github.com/rajaabluu/ershop-api/internal/entity"

type CustomerRepository struct {
	Repository[entity.Customer]
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}
