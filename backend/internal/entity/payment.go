package entity

import "time"

type Payment struct {
	ID          uint      `gorm:"primaryKey"`
	PaymentDate time.Time `gorm:"default:current_timestamp"`
	Status      string    `gorm:"default:'PENDING'"`
}
