package repository

import "github.com/rajaabluu/commerce/backend/internal/entity"

type AddressRepository struct {
	Repository[entity.Address]
}

func NewAddressRepository() *AddressRepository {
	return &AddressRepository{}
}
