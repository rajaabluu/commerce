package mapper

import (
	"github.com/rajaabluu/commerce/backend/internal/entity"
	"github.com/rajaabluu/commerce/backend/internal/model"
)

func ToProductResponse(product *entity.Product) *model.ProductResponse {
	var categoriesResponse []*model.Category
	for _, c := range product.Categories {
		categoriesResponse = append(categoriesResponse, &model.Category{
			ID:   c.ID,
			Name: c.Name,
		})
	}
	var imagesResponse []*model.ProductImage
	for _, img := range product.Images {
		imagesResponse = append(imagesResponse, &model.ProductImage{
			ID:     img.ID,
			Source: img.Source,
		})
	}

	return &model.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Categories:  categoriesResponse,
		Images:      imagesResponse,
	}
}
