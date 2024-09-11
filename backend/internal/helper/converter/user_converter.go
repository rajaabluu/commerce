package converter

import (
	"github.com/rajaabluu/ershop-api/internal/entity"
	"github.com/rajaabluu/ershop-api/internal/model"
)

func ToUserResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}
