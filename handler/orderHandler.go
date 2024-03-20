package handler

import (
	"kominfo-assignment-2/dto"
	"kominfo-assignment-2/service/order_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderService order_service.Service
}

func NewOrderHandler(orderService order_service.Service) *orderHandler {
	return &orderHandler{
		orderService: orderService,
	}
}

func (o *orderHandler) CreateOrderWithItems(ctx *gin.Context) {
	var payload dto.CreateOrderRequestDto

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
			"message": "invalid request json",
		})
		return
	}

	if err := o.orderService.CreateOrderWithItems(payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]string{
		"message": "success",
	})
}
