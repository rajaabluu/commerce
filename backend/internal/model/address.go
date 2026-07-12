package model

type CreateAddressRequest struct {
	RecipientName string `json:"recipient_name" validate:"required"`
	Phone         string `json:"phone" validate:"required"`

	Province      string `json:"province" validate:"required"`
	City          string `json:"city" validate:"required"`
	District      string `json:"district" validate:"required"`
	PostalCode    string `json:"postal_code" validate:"required,number"`
	StreetAddress string `json:"street_address" validate:"required"`
}

type AddressResponse struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	RecipientName string `json:"receipent_name"`
	Phone         string `json:"phone"`

	Province      string `json:"province"`
	City          string `json:"city"`
	District      string `json:"district"`
	PostalCode    string `json:"postal_code"`
	StreetAddress string `json:"street_address"`

	IsDefault bool `json:"is_default"`
}
