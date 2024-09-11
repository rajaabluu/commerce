package repository

import (
	"github.com/rajaabluu/ershop-api/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repository *UserRepository) FindByEmail(tx *gorm.DB, email string, entity *entity.User) error {
	return tx.Where("email = ?", email).Find(entity).Error
}
