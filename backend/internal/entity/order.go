package entity

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type Status string

const (
	PENDING  Status = "PENDING"
	REJECTED Status = "REJECTED"
	APPROVED Status = "APPROVED"
)

func (p *Status) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*p = Status(string(v))
	case string:
		*p = Status(v)
	}
	*p = Status(value.([]byte))
	return nil
}

func (p Status) Value() (driver.Value, error) {
	return string(p), nil
}

type Order struct {
	gorm.Model
	UserID uint
	User   User
	Status Status `gorm:"default:'PENDING';type:varchar(20)"`
}
