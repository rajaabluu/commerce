package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `gorm:"not null"`
	Description string         `gorm:"not null"`
	Price       uint           `gorm:"not null"`
	Stock       uint           `gorm:"not null"`
	Categories  []Category     `gorm:"many2many:product_categories;constraint:OnDelete:CASCADE"`
	Images      []ProductImage `gorm:"constraint:OnDelete:CASCADE"`
}
