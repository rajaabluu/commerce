package model

type CreateOrderRequest struct {
	ShippingAddressID uint              `json:"shipping_address_id" validate:"required"`
	PaymentMethod     string            `json:"payment_method" validate:"required"`
	Items             []CreateOrderItem `json:"items" validate:"required,min=1,dive"`
}

type CreateOrderItem struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required,gt=0"`
}
