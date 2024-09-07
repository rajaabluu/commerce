package entity

import (
	"github.com/rajaabluu/ershop-api/internal/model"
)

type Payment struct {
	ID     uint `gorm:"primaryKey"`
	Status model.Status
}
