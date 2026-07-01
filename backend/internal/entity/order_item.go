package entity

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model

	OrderID uint
	Order   Order

	ProductID uint
	Product   Product

	ProductName string `gorm:"size:255;not null"`
	Price       int64  `gorm:"not null"`
	Quantity    int    `gorm:"not null"`
	Subtotal    int64  `gorm:"not null"`
}
