package entity

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type Status string

const (
	PENDING  Status = "PENDING"
	REJECTED Status = "REJECTED"
	APPROVED Status = "APPROVED"
)

func (p *Status) Scan(value interface{}) error {
	*p = Status(value.([]byte))
	return nil
}

func (p Status) Value() (driver.Value, error) {
	return string(p), nil
}

type Order struct {
	gorm.Model
	UserID uint
	User   User
	Status Status `gorm:"default:'PENDING';type:status"`
}

type OrderDetail struct {
	gorm.Model
	Order     Order
	OrderID   uint
	Product   Product
	ProductID uint
	Quantity  uint
	Price     uint
}
