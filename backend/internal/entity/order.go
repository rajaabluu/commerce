package entity

type Order struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	PaymentID uint
	User      User
	Payment   Payment
}

type OrderDetail struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	ProductID uint
	Price     uint
	Quantity  uint
	Total     uint
	Order     Order
	Product   Product
}
