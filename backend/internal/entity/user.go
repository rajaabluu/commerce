package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Address  string
	Role     uint `gorm:"default:2"`
	Contact  string
}
