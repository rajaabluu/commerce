package entity

type Category struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Products []*Product `gorm:"many2many:product_categories;"`
}
