package model

type CreateCustomerRequest struct {
	Name     string `json:"name,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
	Address  string `json:"address,omitempty" validate:"required"`
	Contact  string `json:"contact,omitempty" validate:"required,number"`
}

type AuthenticateCustomerRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type CustomerResponse struct {
	ID      uint   `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	Address string `json:"address,omitempty"`
	Role    string `json:"role,omitempty"`
}
