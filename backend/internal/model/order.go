package model

type CreateOrderRequest struct {
	UserID    uint
	ProductID uint `json:"product_id,omitempty"`
	Price     uint `json:"price,omitempty"`
	Quantity  uint `json:"quantity"`
}
