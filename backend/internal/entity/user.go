package entity

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Role string

const (
	ADMIN    Role = "ADMIN"
	CUSTOMER Role = "CUSTOMER"
)

func (p *Role) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*p = Role(string(v))
	case string:
		*p = Role(v)
	default:
		return fmt.Errorf("unsupported Scan type for Role: %T", value)
	}
	return nil
}
func (p Role) Value() (driver.Value, error) {
	return string(p), nil
}

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	Contact   *string
	Address   *string
	Role      Role `gorm:"type:role;default:'CUSTOMER'"`
	CreatedAt time.Time
	DeletedAt time.Time
}
