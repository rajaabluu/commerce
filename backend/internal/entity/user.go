package entity

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Address  string
	Role     string `gorm:"default:2"`
	Contact  string
}
