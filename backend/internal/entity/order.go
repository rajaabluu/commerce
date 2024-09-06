package entity

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderID   uint
	UserID    uint
	PaymentID uint
	User      User
	OrderDate time.Time
}
