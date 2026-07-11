package mapper

import (
	"github.com/rajaabluu/commerce/backend/internal/entity"
	"github.com/rajaabluu/commerce/backend/internal/model"
)

func ToUserResponse(user *entity.User) *model.UserResponse {
	phone := ""
	if user.Phone != nil {
		phone = *user.Phone
	}
	return &model.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: phone,
		Role:  string(user.Role),
	}
}
