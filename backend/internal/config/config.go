package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
	Uploader UploaderConfig
	JWT      JWTConfig
	App      AppConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
	SSLMode  string
}

type UploaderConfig struct {
	CloudName string
	APISecret string
	APIKey    string
}

type JWTConfig struct {
	Secret string
}

type AppConfig struct {
	Port    int
	EnvMode string
}

func NewConfig() *Config {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error in reading config: %v", err))
	}

	config := new(Config)

	if err := v.Unmarshal(config); err != nil {
		panic(fmt.Errorf("failed to parse config: %+v", err))
	}

	return config

}
