package routes

import (
	"github.com/coleYab/erestourant/internal/handler"
	"github.com/gin-gonic/gin"
)

type OrderRoutes struct {
	e *gin.Engine
}

func NewOrderRoutes(e *gin.Engine) *OrderRoutes {
	return &OrderRoutes{e: e}
}

func (r *OrderRoutes) RegisterRoutes(menuHandler *handler.OrderHandler) {
	r.e.GET("/order/", menuHandler.GetOrders)
	r.e.GET("/verify/order/:id", menuHandler.ValidateOrder)
	r.e.GET("/order/:id", menuHandler.GetDetails)
	r.e.POST("/order/", menuHandler.CreateOrder)
	r.e.DELETE("/order/:id", menuHandler.DeleteOrder)
	r.e.PUT("/order/:id", menuHandler.EditOrder)
}
