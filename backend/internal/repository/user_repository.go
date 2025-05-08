package repository

import "github.com/rajaabluu/commerce/backend/internal/entity"

type UserRepository struct {
	Repository[entity.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
