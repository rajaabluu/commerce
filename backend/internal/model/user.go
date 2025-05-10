package model

type CreateUserRequest struct {
	Name     string `json:"name,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty"`
	Contact  string `json:"contact,omitempty"`
	Address  string `json:"address,omitempty"`
}

type UserResponse struct {
	ID    uint   `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
}

type UserTokenResponse struct {
	AccessToken string `json:"access_token,omitempty"`
}
