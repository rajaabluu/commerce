package mapper

import (
	"github.com/rajaabluu/commerce/backend/internal/entity"
	"github.com/rajaabluu/commerce/backend/internal/model"
)

func ToAddressEntity(req *model.CreateAddressRequest) *entity.Address {
	return &entity.Address{
		RecipientName: req.RecipientName,
		Phone:         req.Phone,
		Province:      req.Province,
		District:      req.District,
		City:          req.City,
		StreetAddress: req.StreetAddress,
		PostalCode:    req.PostalCode,
	}
}

func ToAddressResponse(address *entity.Address) *model.AddressResponse {
	return &model.AddressResponse{
		ID:            address.ID,
		UserID:        address.UserID,
		RecipientName: address.RecipientName,
		Phone:         address.Phone,
		District:      address.District,
		City:          address.City,
		Province:      address.Province,
		PostalCode:    address.PostalCode,
		StreetAddress: address.StreetAddress,
		IsDefault:     address.IsDefault,
	}
}
