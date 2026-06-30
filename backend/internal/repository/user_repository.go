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

func (r *UserRepository) FindByEmail(db *gorm.DB, email string) (*entity.User, error) {
	user := new(entity.User)
	return user, db.Where("email = ?", email).Take(user).Error
}
