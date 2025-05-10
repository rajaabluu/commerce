package repository

import (
	"github.com/rajaabluu/commerce/backend/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repository *UserRepository) FindByEmail(db *gorm.DB, email string, user *entity.User) error {
	return db.Where("email = ?", email).Find(user).Error
}
