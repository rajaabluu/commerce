package entity

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	Contact   string
	Address   *string
	Role 	  uint `gorm:"default:2"`
	CreatedAt time.Time
	DeletedAt time.Time
}

