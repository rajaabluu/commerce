package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rajaabluu/commerce/backend/internal/model"
	"github.com/spf13/viper"
)

func GenerateToken(config *viper.Viper, user *model.UserResponse) (string, error) {
	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": jwt.NewNumericDate(time.Now().Add(24 * 30 * time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetString("jwt.secret")))
}
