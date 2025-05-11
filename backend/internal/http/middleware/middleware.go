package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rajaabluu/commerce/backend/internal/config"
	"github.com/rajaabluu/commerce/backend/internal/helper"
	"github.com/rajaabluu/commerce/backend/internal/model"
	"github.com/sirupsen/logrus"
)

type Middleware struct {
	Logger *logrus.Logger
	Config *config.Config
}

func NewMiddleware(logger *logrus.Logger, config *config.Config) *Middleware {
	return &Middleware{logger, config}
}

func (m *Middleware) VerifyAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		schema := "Bearer "
		authHeader := c.Request().Header.Get("Authorization")
		if len(authHeader) == 0 {
			return c.JSON(http.StatusUnauthorized, &model.ErrorResponse{
				Message: "unauthorized user",
			})
		}
		tokenString := authHeader[len(schema):]
		claims, err := helper.ValidateToken(m.Config, tokenString)
		if err != nil {
			m.Logger.Warnf("error on validating token: %+v", err)
			return c.JSON(http.StatusUnauthorized, &model.ErrorResponse{
				Message: "unauthorized user",
			})
		}

		c.Set("userId", claims["id"])
		c.Set("role", claims["role"])

		return next(c)
	}
}
