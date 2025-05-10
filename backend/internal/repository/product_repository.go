package repository

import "github.com/rajaabluu/commerce/backend/internal/entity"

type ProductRepository struct {
	Repository[entity.Product]
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}
