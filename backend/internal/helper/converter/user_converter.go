package converter

import (
	"github.com/rajaabluu/ershop-api/internal/entity"
	"github.com/rajaabluu/ershop-api/internal/model"
)

func ToCustomerResponse(Customer *entity.Customer) *model.CustomerResponse {
	return &model.CustomerResponse{
		ID:      Customer.ID,
		Name:    Customer.Name,
		Email:   Customer.Email,
		Address: Customer.Address,
	}
}
