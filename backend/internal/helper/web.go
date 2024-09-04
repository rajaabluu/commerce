package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rajaabluu/ershop-api/internal/model"
	"github.com/spf13/viper"
)

func WriteJSONResponse(w http.ResponseWriter, response interface{}, status int) {
	if status == 0 {
		status = http.StatusOK
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatalf("error on writing json response: %v", err)
	}
}

func DecodeRequestBody(r *http.Request, toDecode interface{}) error {
	return json.NewDecoder(r.Body).Decode(toDecode)
}

func GenerateToken(viper *viper.Viper, user *model.CustomerResponse) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	return t.SignedString([]byte(viper.GetString("jwt.secret")))
}

func ValidateToken(viper *viper.Viper, tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("jwt.secret")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, errors.New("error on parsing token, user unauthorized")
	}
}
