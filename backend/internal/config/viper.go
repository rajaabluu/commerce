package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(".")
	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error in reading config: %w", err))
	}
	return config
}
