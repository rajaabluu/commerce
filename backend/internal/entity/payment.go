package entity

import (
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

	OrderID uint `gorm:"uniqueIndex"`
	Order   Order

	Status PaymentStatus `gorm:"type:varchar(20);default:'PENDING'"`

	Method string `gorm:"size:30"`

	TransactionID string
}
