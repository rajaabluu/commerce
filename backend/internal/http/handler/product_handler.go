package handler

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rajaabluu/commerce/backend/internal/helper"
	"github.com/rajaabluu/commerce/backend/internal/model"
	"github.com/rajaabluu/commerce/backend/internal/service"
	"github.com/sirupsen/logrus"
)

type ProductHandler struct {
	ProductService *service.ProductService
	Logger         *logrus.Logger
}

func NewProductHandler(logger *logrus.Logger, productService *service.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: productService,
		Logger:         logger,
	}
}

func (handler *ProductHandler) GetProduct(c echo.Context) error {
	return errors.New("not implemented yet")
}

func (handler *ProductHandler) CreateNewProduct(c echo.Context) error {
	req := new(model.CreateProductRequest)

	if err := c.Bind(req); err != nil {
		handler.Logger.Warnf("error on decoding request: %+v", err)
		return err
	}

	res, err := handler.ProductService.Create(c.Request().Context(), req)

	if err != nil {
		var ve validator.ValidationErrors
		switch {
		case errors.As(err, &ve):
			return c.JSON(http.StatusUnprocessableEntity, &model.ErrorResponse{
				Message: "validation errors",
				Error:   helper.GenerateValidationError(ve),
			})
		}
		return c.JSON(http.StatusInternalServerError, &model.ErrorResponse{
			Message: "internal server error",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &model.Response[*model.ProductResponse]{
		Message: "product successfully created",
		Data:    res,
	})
}

func (handler *ProductHandler) GetProductById(c echo.Context) error {
	return errors.New("not implemented yet")
}

func (handler *ProductHandler) DeleteProduct(c echo.Context) error {
	return errors.New("not implemented yet")
}
