package model

type CreateUserRequest struct {
	Name     string `json:"name,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=6"`
	Contact  string `json:"contact,omitempty"`
	Address  string `json:"address,omitempty"`
}

type AuthenticateUserRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserResponse struct {
	ID    uint   `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
	TokenResponse
}

type UpdateUserProfileRequest struct {
	ID      uint    `json:"id,omitempty"`
	Name    *string `json:"name"`
	Email   *string `json:"email"`
	Contact *string `json:"contact"`
	Address *string `json:"address"`
	Role    string  `json:"role,omitempty"`
}

type UserProfileResponse struct {
	ID      uint    `json:"id,omitempty"`
	Name    string  `json:"name,omitempty"`
	Email   string  `json:"email,omitempty"`
	Contact *string `json:"contact"`
	Address *string `json:"address"`
	Role    string  `json:"role,omitempty"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token,omitempty"`
}
