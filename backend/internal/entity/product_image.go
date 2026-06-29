package entity

type ProductImage struct {
	ID        uint `gorm:"primaryKey"`
	ProductID uint
	Source    string
	PublicID  string
}
