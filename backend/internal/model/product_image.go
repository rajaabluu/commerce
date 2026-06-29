package model

type ProductImage struct {
	ID        uint   `json:"id,omitempty"`
	ProductID uint   `json:"product_id,omitempty"`
	Source    string `json:"source,omitempty"`
	PublicID  string `json:"public_id,omitempty"`
}

type ProductImageResponse = ProductImage
