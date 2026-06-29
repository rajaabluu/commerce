package repository

import (
	"github.com/rajaabluu/commerce/backend/internal/entity"
	"gorm.io/gorm"
)

type ProductImageRepository struct {
	Repository[entity.ProductImage]
}

func NewProductImageRepository() *ProductImageRepository {
	return &ProductImageRepository{}
}

func (r *ProductImageRepository) BulkCreate(db *gorm.DB, entities []*entity.ProductImage) error {
	return db.Create(entities).Error
}
