package config

import (
	"fmt"

	"github.com/rajaabluu/commerce/backend/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(config *Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Database.Host,
		config.Database.User,
		config.Database.Password,
		config.Database.DBName,
		config.Database.Port,
		config.Database.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// db.Exec(`CREATE TYPE role AS ENUM ('ADMIN', 'CUSTOMER')`)
	// db.Exec(`CREATE TYPE status AS ENUM ('APPROVED', 'PENDING', 'REJECTED')`)

	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Product{},
		&entity.Category{},
		&entity.Order{},
		&entity.Payment{},
		&entity.OrderDetail{},
	); err != nil {
		panic(fmt.Errorf("failed migrating database: %w", err))
	}

	if err != nil {
		panic(fmt.Errorf("error in connecting to database: %w", err))
	}

	return db
}
