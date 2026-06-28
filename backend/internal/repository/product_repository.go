package repository

import (
	"github.com/rajaabluu/commerce/backend/internal/entity"
	"github.com/rajaabluu/commerce/backend/internal/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Repository[entity.Product]
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) GetProductsByFilter(tx *gorm.DB, filter *model.ProductFilter)

func (r *ProductRepository) AddProductImage(tx *gorm.DB, img *entity.ProductImage) error {
	return tx.Create(img).Error
}
