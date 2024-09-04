package config

import (
	"log"

	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("./")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	return config
}
