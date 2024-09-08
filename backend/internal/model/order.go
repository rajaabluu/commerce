package model

type CreateOrderRequest struct {
	Products []*CreateOrderProductRequest
}

type CreateOrderResponse struct {
	RedirectUrl string `json:"redirect_url,omitempty"`
}

type OrderResponse struct {
	ID       string             `json:"id,omitempty"`
	UserID   uint               `json:"user_id,omitempty"`
	User     UserResponse       `json:"user,omitempty"`
	Products []*ProductResponse `json:"products,omitempty"`
	Total    int                `json:"total,omitempty"`
}
