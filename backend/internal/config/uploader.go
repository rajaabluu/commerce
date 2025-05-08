package config

import (
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/spf13/viper"
)

func NewUploader(*viper.Viper) *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromParams(
		viper.GetString("cloudinary.cloud.name"),
		viper.GetString("cloudinary.api.key"),
		viper.GetString("cloudinary.api.secret"))

	if err != nil {
		panic(fmt.Errorf("error on connecting cloudinary: %w", err))
	}

	return cld
}
