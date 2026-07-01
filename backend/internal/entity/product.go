package entity

import "time"

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Price       uint
	Stock       uint
	Categories  []Category `gorm:"many2many:product_categories;constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time
	DeletedAt   time.Time
	Images      []ProductImage `gorm:"constraint:OnDelete:CASCADE"`
}
