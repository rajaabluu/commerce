package entity

import (
	"time"
)

type Product struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Slug        string
	Description string
	Quantity    uint
	Price       uint
	Categories  []*Category     `gorm:"many2many:product_categories;constraint:OnDelete:CASCADE"`
	Images      []*ProductImage `gorm:"constraint:OnDelete:CASCADE"`
}

type ProductImage struct {
	ID        uint `gorm:"primarykey"`
	ProductID uint
	Source    string
	PublicID  string
}
