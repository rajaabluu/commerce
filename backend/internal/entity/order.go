package entity

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderID    uint
	CustomerID uint
	PaymentID  uint
	Customer   Customer
	OrderDate  time.Time
}
