package config

import (
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
)

func NewUploader(config *Config) *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromParams(
		config.Uploader.CloudName,
		config.Uploader.APIKey,
		config.Uploader.APISecret)

	if err != nil {
		panic(fmt.Errorf("error on connecting cloudinary: %w", err))
	}

	return cld
}
