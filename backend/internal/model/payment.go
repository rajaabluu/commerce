package model

type CreatePaymentResponse struct {
	RedirectUrl string `json:"redirect_url"`
}

type CreatePaymentRequest struct {
	Products []*CreatePaymentProductRequest
	OrderID  string
}
