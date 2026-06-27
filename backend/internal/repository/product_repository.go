package repository

import (
	"github.com/rajaabluu/commerce/backend/internal/entity"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Repository[entity.Product]
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (repository *ProductRepository) AddProductImage(tx *gorm.DB, img *entity.ProductImage) error {
	return tx.Create(img).Error
}
