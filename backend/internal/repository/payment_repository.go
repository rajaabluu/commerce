package repository

import (
	"github.com/rajaabluu/commerce/backend/internal/entity"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	Repository[entity.Payment]
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{}
}

func (r *PaymentRepository) FindOne(db *gorm.DB, conds *entity.Payment) (*entity.Payment, error) {
	payment := new(entity.Payment)
	err := db.Where(conds).First(payment).Error
	return payment, err
}
