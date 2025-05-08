package entity

import "time"

type Payment struct {
	ID        uint
	CreatedAt time.Time
	DeletedAt time.Time
}
