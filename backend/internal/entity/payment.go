package entity

import (
	"time"

	"gorm.io/gorm"
)

type PaymentStatus string

const (
	PaymentPending PaymentStatus = "PENDING"
	PaymentFailed  PaymentStatus = "FAILED"
	PaymentPaid    PaymentStatus = "PAID"
	PaymentExpired PaymentStatus = "EXPIRED"
)

type Payment struct {
	gorm.Model
	ID            uint
	OrderID       uint
	Order         Order
	Status        PaymentStatus
	Method        string
	TransactionID string
	CreatedAt     time.Time
	DeletedAt     time.Time
}
