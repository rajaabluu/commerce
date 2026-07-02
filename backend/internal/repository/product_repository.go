package repository

import (
	"strings"

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

func (r *ProductRepository) Find(db *gorm.DB, filter *model.ProductFilter) ([]*entity.Product, error) {

	var products []*entity.Product
	if filter.Search != "" {
		db = db.Where("name ILIKE ?", "%"+filter.Search+"%")
	}

	if len(filter.Categories) > 0 {
		db = db.
			Joins("JOIN product_categories pc ON pc.product_id = products.id").
			Where("pc.category_id IN ?", filter.Categories).
			Distinct("products.*")
	}

	if filter.SortBy != "" {
		allowedSort := map[string]string{
			"name":       "name",
			"price":      "price",
			"created_at": "created_at",
		}

		column, ok := allowedSort[filter.SortBy]
		if !ok {
			column = "created_at"
		}

		order := "DESC"

		if strings.EqualFold(filter.SortOrder, "asc") {
			order = "ASC"
		}

		db = db.Order(column + " " + order)
	}

	if filter.MinPrice > 0 {
		db = db.Where("price >= ?", filter.MinPrice)
	}

	if filter.MaxPrice > 0 {
		db = db.Where("price <= ?", filter.MaxPrice)
	}

	if filter.Limit > 0 {
		db = db.Limit(filter.Limit).Offset(filter.Offset)
	}

	err := db.Preload("Categories").Preload("Images").Find(&products).Error

	return products, err
}

func (r *ProductRepository) Update(db *gorm.DB, product *entity.Product) error {
	if err := db.
		Model(&entity.Product{}).
		Where("id = ?", product.ID).
		Updates(product).Error; err != nil {
		return err
	}

	return db.
		Preload("Images").
		Preload("Categories").
		First(product, product.ID).Error
}

func (r *ProductRepository) FindByIDs(
	db *gorm.DB,
	ids []uint,
) ([]*entity.Product, error) {
	var products []*entity.Product

	if err := db.
		Where("id IN ?", ids).
		Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
