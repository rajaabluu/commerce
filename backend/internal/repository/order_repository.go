package repository

import "github.com/rajaabluu/commerce/backend/internal/entity"

type OrderRepository struct {
	Repository[entity.Order]
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}
