package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string `gorm:"default:null"`
	Address  string `gorm:"default:null"`
	Role     uint   `gorm:"default:2"`
}
