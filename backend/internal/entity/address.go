package entity

import "gorm.io/gorm"

type Address struct {
	gorm.Model

	UserID        uint   `gorm:"not null"`
	RecipientName string `gorm:"not null"`
	Phone         string `gorm:"size:20;not null"`

	Province      string `gorm:"not null"`
	City          string `gorm:"not null"`
	District      string `gorm:"not null"`
	PostalCode    string `gorm:"not null"`
	StreetAddress string `gorm:"type:text;not null"`

	IsDefault bool `gorm:"default:false"`
}
