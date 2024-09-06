package config

import (
	"fmt"
	"log"

	"github.com/rajaabluu/ershop-api/internal/entity"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(viper *viper.Viper) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", viper.GetString("database.host"), viper.GetString("database.user"), viper.GetString("database.password"), viper.GetString("database.name"), viper.GetInt("database.port"))
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	database.AutoMigrate(&entity.User{}, &entity.Product{}, &entity.Category{}, &entity.ProductImage{})
	return database
}
