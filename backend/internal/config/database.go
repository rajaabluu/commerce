package config

import (
	"fmt"

	"github.com/rajaabluu/commerce/backend/internal/entity"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(viper *viper.Viper) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		viper.GetString("database.host"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.name"),
		viper.GetInt("database.port"),
		viper.GetString("database.ssl"))

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
