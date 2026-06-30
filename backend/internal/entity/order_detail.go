package entity

import "gorm.io/gorm"

type OrderDetail struct {
	gorm.Model
	Order     Order
	OrderID   uint
	Product   Product
	ProductID uint
	Quantity  uint
	Price     uint
}
