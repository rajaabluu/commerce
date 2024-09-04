package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/spf13/viper"
)

func NewUploader(viper *viper.Viper) *cloudinary.Cloudinary {
	// url := fmt.Sprintf("cloudinary://%s:%s@%s", viper.GetString("cloudinary.api.key"), viper.GetString("cloudinary.api.secret"), viper.GetString("cloudinary.cloud.name"))
	cld, _ := cloudinary.NewFromParams(viper.GetString("cloudinary.cloud.name"), viper.GetString("cloudinary.api.key"), viper.GetString("cloudinary.api.secret"))
	cld.Config.URL.Secure = true
	return cld
}
