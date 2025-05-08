package entity

import "time"

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Price       uint
	Stock       uint
	CategoryID  uint
	Categories  []Category `gorm:"many2many:product_categories;constraint;OnDelete:CASCADE"`
	CreatedAt   time.Time
	DeletedAt   time.Time
}

type ProductImage struct {
	ID        uint `gorm:"primaryKey"`
	ProductID uint
	Source    string
	PublicID  string
}
