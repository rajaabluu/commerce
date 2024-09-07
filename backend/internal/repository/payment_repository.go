package repository

import "github.com/rajaabluu/ershop-api/internal/entity"

type PaymentRepository struct {
	Repository[entity.Payment]
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{}
}
