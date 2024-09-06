package repository

import "github.com/rajaabluu/ershop-api/internal/entity"

type UserRepository struct {
	Repository[entity.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
