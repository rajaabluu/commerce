package converter

import (
	"github.com/rajaabluu/ershop-api/internal/entity"
	"github.com/rajaabluu/ershop-api/internal/model"
)

func ToProductResponse(product *entity.Product) *model.ProductResponse {
	var categories []*model.Category
	var images []*model.ProductImage
	for _, val := range product.Categories {
		categories = append(categories, &model.Category{Name: val.Name, ID: val.ID})
	}
	for _, val := range product.Images {
		images = append(images, &model.ProductImage{ProductID: val.ProductID, Source: val.Source, PublicID: val.PublicID})
	}
	return &model.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Quantity:    product.Quantity,
		Price:       product.Price,
		Categories:  categories,
		Images:      images,
	}
}

// func ToProductRequest() *model.ProductRequest {
// 	/// erro

// }
