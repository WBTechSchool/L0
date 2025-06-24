package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/WBTechSchool/L0/internal/service"
	"github.com/gin-gonic/gin"
)

type OrderHandlers interface {
	GetOrderByID(ctx *gin.Context)
}

type orderHandlers struct {
	s *service.Service
}

func NewOrderHandlers(s *service.Service) OrderHandlers {
	return &orderHandlers{
		s: s,
	}
}

func (o *orderHandlers) GetOrderByID(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(JSONErrorWrapper(http.StatusBadRequest, "Invalid order ID"))
		return
	}

	order, err := o.s.OrderService.GetOrderByID(context.Background(), id)
	if err != nil {
		ctx.JSON(JSONErrorWrapper(http.StatusInternalServerError, "Order not found"))
		return
	}

	ctx.JSON(http.StatusOK, order)
}
