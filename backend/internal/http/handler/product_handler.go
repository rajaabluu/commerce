package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/rajaabluu/ershop-api/internal/helper"
	"github.com/rajaabluu/ershop-api/internal/model"
	"github.com/rajaabluu/ershop-api/internal/service"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductHandler struct {
	ProductService *service.ProductService
	Logger         *logrus.Logger
}

func NewProductHandler(logger *logrus.Logger, service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: service,
		Logger:         logger,
	}
}

func (handler *ProductHandler) GetProductCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := handler.ProductService.GetCategories(r.Context())
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			helper.WriteJSONResponse(w, &model.Response[any]{Message: "products retrieved", Data: &[]*model.ProductResponse{}}, http.StatusOK)
			return
		default:
			helper.WriteJSONResponse(w, &model.ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
	}
	helper.WriteJSONResponse(w, &model.Response[[]*model.Category]{Message: "success get all categories", Data: categories}, http.StatusOK)
}

func (handler *ProductHandler) CreateNewProduct(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(4096); err != nil {
		helper.WriteJSONResponse(w, &model.ErrResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}

	qty, _ := strconv.ParseUint(strings.TrimSpace(r.FormValue("quantity")), 10, 32)
	price, _ := strconv.ParseUint(strings.TrimSpace(r.FormValue("price")), 10, 32)

	handler.Logger.Warn(qty, price)

	product := &model.CreateProductRequest{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Quantity:    uint(qty),
		Price:       uint(price),
		Categories:  r.Form["categories[]"],
		Images:      r.MultipartForm.File["images[]"],
	}
	result, err := handler.ProductService.CreateProduct(r.Context(), product)
	if err != nil {
		helper.WriteJSONResponse(w, &model.ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	helper.WriteJSONResponse(w, &model.Response[*model.ProductResponse]{Message: "success creating product", Data: result}, http.StatusCreated)
}

func (handler *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var products []*model.ProductResponse
	var err error
	query := r.URL.Query()["categories[]"]
	if len(query) > 0 {
		products, err = handler.ProductService.GetProductsByCategory(r.Context(), query)
	} else {
		products, err = handler.ProductService.GetAllProducts(r.Context())
	}
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			helper.WriteJSONResponse(w, &model.Response[any]{Message: "products retrieved", Data: &[]*model.ProductResponse{}}, http.StatusOK)
			return
		default:
			helper.WriteJSONResponse(w, &model.ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
	}
	helper.WriteJSONResponse(w, &model.Response[*[]*model.ProductResponse]{
		Message: "products retrieved",
		Data:    &products,
	}, http.StatusOK)
}

func (handler *ProductHandler) GetProductDetail(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	product, err := handler.ProductService.GetProduct(r.Context(), uint(id))
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			helper.WriteJSONResponse(w, &model.ErrResponse{Message: "product not found"}, http.StatusNotFound)
			return
		default:
			helper.WriteJSONResponse(w, &model.ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
	}
	helper.WriteJSONResponse(w, &model.Response[*model.ProductResponse]{
		Message: "product retrieved",
		Data:    product,
	}, http.StatusOK)
}

func (handler *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
	if err := handler.ProductService.DeleteProduct(r.Context(), uint(id)); err != nil {
		helper.WriteJSONResponse(w, &model.ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	helper.WriteJSONResponse(w, &model.Response[any]{Message: "Product Deleted"}, http.StatusOK)
}
