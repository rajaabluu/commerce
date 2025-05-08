package entity

import (
	"gorm.io/gorm"
)

type Status string

const (
	PENDING  Status = "PENDING"
	REJECTED Status = "REJECTED"
	APPROVED Status = "APPROVED"
)

type Order struct {
	gorm.Model
	UserID uint
	User   User
	Status Status `gorm:"default:'PENDING'"`
}
