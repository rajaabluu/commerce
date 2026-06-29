package handler

import (
	"errors"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rajaabluu/commerce/backend/internal/exception"
	"github.com/rajaabluu/commerce/backend/internal/model"
	"github.com/rajaabluu/commerce/backend/internal/service"
	"github.com/sirupsen/logrus"
)

type ProductImageHandler struct {
	Logger              *logrus.Logger
	ProductImageService *service.ProductImageService
}

func NewProductImageHandler(logger *logrus.Logger, productImgService *service.ProductImageService) *ProductImageHandler {
	return &ProductImageHandler{
		Logger:              logger,
		ProductImageService: productImgService,
	}
}

func (h *ProductImageHandler) Upload(c echo.Context) error {
	var productID uint
	form, err := c.MultipartForm()
	var files []*multipart.FileHeader

	if param := c.Param("id"); param != "" {
		p, err := strconv.Atoi(param)
		if err != nil {
			h.Logger.Warnf("invalid param type")
			return c.JSON(http.StatusBadRequest, &model.ErrorResponse{
				Message: "invalid product id",
			})
		}
		productID = uint(p)
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, &model.ErrorResponse{
			Message: "invalid file",
		})
	}

	for _, file := range form.File["image"] {
		files = append(files, file)
	}

	res, err := h.ProductImageService.UploadProductImage(c.Request().Context(), files, productID)

	if err != nil {
		switch {
		case errors.Is(err, exception.ErrNoProductImages):
			return c.JSON(http.StatusBadRequest, &model.ErrorResponse{
				Message: "no image specified",
			})
		case errors.Is(err, exception.ErrTooManyProductImages):
			return c.JSON(http.StatusBadRequest, &model.ErrorResponse{
				Message: "too many images. max 5",
			})
		default:
			return c.JSON(http.StatusInternalServerError, &model.ErrorResponse{
				Message: "error on processing",
			})
		}

	}

	return c.JSON(http.StatusOK, &model.Response[[]*model.ProductImageResponse]{
		Message: "product images uploaded sucessfully",
		Data:    res,
	})
}
