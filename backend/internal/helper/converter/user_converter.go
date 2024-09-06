package converter

import (
	"github.com/rajaabluu/ershop-api/internal/entity"
	"github.com/rajaabluu/ershop-api/internal/model"
)

func ToUserResponse(customer *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:    customer.ID,
		Name:  customer.Name,
		Email: customer.Email,
		Role:  customer.Role,
	}
}
