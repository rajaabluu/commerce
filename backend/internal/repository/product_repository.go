package repository

import "github.com/rajaabluu/ershop-api/internal/entity"

type ProductRepository struct {
	Repository[entity.Product]
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}
