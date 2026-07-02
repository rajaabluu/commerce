package model

type CreateOrderRequest struct {
	ShippingAddressID uint              `json:"shipping_address_id" validate:"required"`
	PaymentMethod     string            `json:"payment_method" validate:"required"`
	OrderItems        []CreateOrderItem `json:"items" validate:"required,min=1,dive"`
}

type CreateOrderItem struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required,gt=0"`
}

type OrderItemResponse struct {
	ID          uint   `json:"id"`
	ProductID   uint   `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       int64  `json:"price"`
	Quantity    uint   `json:"quantity"`
	Subtotal    int64  `json:"subtotal"`
}

type OrderResponse struct {
	ID uint `json:"id"`

	Status string `json:"status"`

	PaymentMethod string `json:"payment_method"`

	TotalPrice int64 `json:"total_price"`

	RecipientName string `json:"recipient_name"`
	Phone         string `json:"phone"`

	Province      string `json:"province"`
	City          string `json:"city"`
	District      string `json:"district"`
	PostalCode    string `json:"postal_code"`
	StreetAddress string `json:"street_address"`

	OrderItems []*OrderItemResponse `json:"items"`
}

type CreateOrderResponse = OrderResponse
