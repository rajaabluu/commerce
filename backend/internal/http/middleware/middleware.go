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

func (m *Middleware) VerifyIsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Get("role")
		if role == nil {
			return c.JSON(http.StatusForbidden, &model.ErrorResponse{
				Message: "forbidden",
				Error:   "role not found",
			})
		}

		if role != "ADMIN" {
			return c.JSON(http.StatusForbidden, &model.ErrorResponse{
				Message: "forbidden",
				Error:   "admin access required",
			})
		}

		return next(c)
	}
}
