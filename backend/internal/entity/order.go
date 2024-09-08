package entity

import "time"

type Order struct {
	ID        string `gorm:"primaryKey"`
	UserID    uint
	PaymentID *uint
	User      User
	Payment   Payment
	OrderDate time.Time `gorm:"default:current_timestamp"`
}

type OrderDetail struct {
	ID        uint `gorm:"primaryKey,autoIncrement"`
	OrderID   string
	ProductID uint
	Quantity  uint
	Product   Product
	Total     uint
}
