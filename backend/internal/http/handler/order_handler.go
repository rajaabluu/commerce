package handler

import (
	"net/http"

	"github.com/rajaabluu/ershop-api/internal/helper"
	"github.com/rajaabluu/ershop-api/internal/model"
	"github.com/rajaabluu/ershop-api/internal/service"
)

type OrderHandler struct {
	OrderService *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{OrderService: service}
}

func (handler *OrderHandler) NewOrder(w http.ResponseWriter, r *http.Request) {
	request := new(model.CreateOrderRequest)
	if err := helper.DecodeRequestBody(r, request); err != nil {
		helper.WriteJSONResponse(w, &model.ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	response, err := handler.OrderService.CreateOrder(r.Context(), request)
	if err != nil {
		helper.WriteJSONResponse(w, &model.ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	helper.WriteJSONResponse(w, &model.Response[*model.CreateOrderResponse]{Message: "success creating order", Data: response}, http.StatusCreated)
}
