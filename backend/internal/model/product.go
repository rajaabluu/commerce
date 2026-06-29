package model

type CreateProductRequest struct {
	Name        string `json:"name,omitempty" validate:"required"`
	Description string `json:"description,omitempty" validate:"required"`
	Price       uint   `json:"price,omitempty" validate:"required"`
	Stock       uint   `json:"stock,omitempty" validate:"required"`
	CategoryIds []int  `json:"category_ids,omitempty" validate:"required"`
}

type ProductResponse struct {
	ID          uint            `json:"id,omitempty"`
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Price       uint            `json:"price,omitempty"`
	Stock       uint            `json:"stock,omitempty"`
	Categories  []*Category     `json:"categories,omitempty"`
	Images      []*ProductImage `json:"images"`
}

type Category struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type GetProductsRequest struct {
	Page       int
	Limit      int
	Search     string
	Categories []string

	MinPrice float64
	MaxPrice float64

	SortBy    string
	SortOrder string
}

type ProductFilter struct {
	Limit  int
	Offset int

	Search     string
	Categories []string

	MinPrice float64
	MaxPrice float64

	SortBy    string
	SortOrder string
}
