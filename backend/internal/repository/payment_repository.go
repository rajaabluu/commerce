package repository

import "github.com/rajaabluu/commerce/backend/internal/entity"

type PaymentRepository struct {
	Repository[entity.Payment]
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{}
}
