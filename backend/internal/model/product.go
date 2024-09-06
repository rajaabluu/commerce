package model

import "mime/multipart"

type Category struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProductResponse struct {
	ID          uint            `json:"id,omitempty"`
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Quantity    uint            `json:"quantity,omitempty"`
	Price       uint            `json:"price,omitempty"`
	Categories  []*Category     `json:"categories,omitempty"`
	Images      []*ProductImage `json:"images,omitempty"`
}

type CreateProductRequest struct {
	Name        string                  `json:"name,omitempty"`
	Description string                  `json:"description"`
	Quantity    uint                    `json:"quantity,omitempty"`
	Price       uint                    `json:"price,omitempty"`
	Categories  []string                `json:"categories,omitempty"`
	Images      []*multipart.FileHeader `json:"images,omitempty"`
}

type EditProductRequest struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description"`
	Quantity    uint     `json:"quantity,omitempty"`
	Price       uint     `json:"price,omitempty"`
	Categories  []string `json:"categories,omitempty"`
}

type ProductImage struct {
	ProductID uint   `json:"-"`
	Source    string `json:"source,omitempty"`
	PublicID  string `json:"public_id,omitempty"`
}
