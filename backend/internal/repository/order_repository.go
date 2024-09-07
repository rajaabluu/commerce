package repository

import "github.com/rajaabluu/ershop-api/internal/entity"

type OrderRepository struct {
	Repository[entity.Order]
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}
