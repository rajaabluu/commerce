package model

import "github.com/rajaabluu/commerce/backend/internal/entity"

type CreateProductRequest struct {
	Name        string `json:"name,omitempty" validate:"required"`
	Description string `json:"description,omitempty" validate:"required"`
	Price       uint   `json:"price,omitempty" validate:"required"`
	Stock       uint   `json:"stock,omitempty" validate:"required"`
	CategoryIds []int  `json:"category_ids,omitempty" validate:"required"`
}

type ProductResponse struct {
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	Price       uint               `json:"price,omitempty"`
	Stock       uint               `json:"stock,omitempty"`
	Categories  *[]entity.Category `json:"categories,omitempty"`
}
