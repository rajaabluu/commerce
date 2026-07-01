package entity

import (
	"database/sql/driver"
	"fmt"

	"gorm.io/gorm"
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
	gorm.Model
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Phone     *string
	Addresses []Address
	Role      Role `gorm:"type:varchar(50);default:'CUSTOMER'"`
}
