package mapper

import (
	"github.com/rajaabluu/commerce/backend/internal/entity"
	"github.com/rajaabluu/commerce/backend/internal/model"
)

func ToAddressEntity(req *model.CreateAddressRequest, userID uint) *entity.Address {
	return &entity.Address{
		UserID:        userID,
		RecipientName: req.RecipientName,
		Phone:         req.Phone,
		Province:      req.Province,
		District:      req.District,
		City:          req.City,
		StreetAddress: req.StreetAddress,
		PostalCode:    req.PostalCode,
	}
}
