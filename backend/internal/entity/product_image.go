package entity

type ProductImage struct {
	ID        uint   `gorm:"primaryKey"`
	ProductID uint   `gorm:"not null"`
	Source    string `gorm:"not null"`
	PublicID  string `gorm:"not null"`
}
