package handler

import (
	"errors"
	"net/http"
	"strconv"

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

func (h *ProductHandler) GetAll(c echo.Context) error {

	req := new(model.GetProductsRequest)

	if page := c.QueryParam("page"); page != "" {
		p, err := strconv.Atoi(page)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid page")
		}
		req.Page = p
	}

	if limit := c.QueryParam("limit"); limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid limit")
		}
		req.Limit = l
	}

	req.Search = c.QueryParam("search")
	req.SortBy = c.QueryParam("sort_by")
	req.SortOrder = c.QueryParam("order")
	req.Categories = c.QueryParams()["category"]

	res, err := h.ProductService.GetProducts(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &model.Response[[]*model.ProductResponse]{
		Message: "products data retrieved",
		Data:    res,
	})

}

func (h *ProductHandler) CreateNewProduct(c echo.Context) error {
	req := new(model.CreateProductRequest)

	if err := c.Bind(req); err != nil {
		h.Logger.Warnf("error on decoding request: %+v", err)
		return err
	}

	res, err := h.ProductService.Create(c.Request().Context(), req)

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

func (h *ProductHandler) GetProductById(c echo.Context) error {
	return errors.New("not implemented yet")
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	return errors.New("not implemented yet")
}
