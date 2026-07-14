package entity

import (
	"gorm.io/gorm"
)

// type OrderStatus string

// const (
// 	OrderPending   OrderStatus = "PENDING"
// 	OrderPaid      OrderStatus = "PAID"
// 	OrderShipped   OrderStatus = "SHIPPED"
// 	OrderCompleted OrderStatus = "COMPLETED"
// 	OrderCancelled OrderStatus = "CANCELLED"
// )

type Order struct {
	gorm.Model
	InvoiceID string `gorm:"unique"`

	UserID uint `gorm:"not null"`
	User   User

	TotalPrice int64 `gorm:"not null"`

	RecipientName string `gorm:"not null"`
	Phone         string `gorm:"not null"`

	Province      string `gorm:"not null"`
	City          string `gorm:"not null"`
	District      string `gorm:"not null"`
	PostalCode    string `gorm:"not null"`
	StreetAddress string `gorm:"type:text;not null"`

	Payment    Payment
	OrderItems []OrderItem
}
