package handler

import (
	"log"
	"net/http"

	"github.com/coleYab/erestourant/internal/db/repository"
	"github.com/coleYab/erestourant/internal/dto"
	"github.com/coleYab/erestourant/internal/store"
	"github.com/coleYab/erestourant/internal/utils"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	st  *store.OrderStore
	mst *store.MenuStore
}

func NewOrderHandler(st *store.OrderStore, mst *store.MenuStore) *OrderHandler {
	return &OrderHandler{st: st, mst: mst}
}

func (h *OrderHandler) GetOrders(ctx *gin.Context) {
	orders, err := h.st.GetAllOrder()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to get all order", "error": err.Error()})
		return
	}
	if orders == nil {
		orders = []repository.Order{}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "orders": orders})
}

func (h *OrderHandler) CreateOrder(ctx *gin.Context) {
	payload := dto.CreateOrderDto{}
	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to parse the order data", "error": err.Error()})
		return
	}

	tot, err := utils.GetTotalOrders(payload, h.mst)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to create the order", "error": err.Error()})
		return
	}

	order, err := h.st.CreateOrder(payload, tot, "6238c716-641b-40ea-bbea-6f3bdb3d2754")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to create the order", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created a order", "order": order})
}

func (h *OrderHandler) EditOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{"message": "editing the order " + id})
}

func (h *OrderHandler) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.st.DeleteOrder(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to delete the order", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "delete the order " + id})
}

func (h *OrderHandler) GetDetails(ctx *gin.Context) {
	id := ctx.Param("id")
	order, err := h.st.GetOrderByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to get the order", "error": err.Error()})
		return
	}

	orderItems, err := h.st.GetOrderItemsByOrderId(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to get the order", "error": err.Error()})
		return
	}

	castedOrder := dto.NewOrderWithDetails(order, orderItems)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"order":   castedOrder,
	})
}

func (h *OrderHandler) ValidateOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	log.Println("This is some error for this shit")

	if err := h.st.VerifyOrder(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to get the order", "error": err.Error()})
		return
	}

	order, err := h.st.GetOrderByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to get the order", "error": err.Error()})
		return
	}

	orderItems, err := h.st.GetOrderItemsByOrderId(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "unable to get the order", "error": err.Error()})
		return
	}

	castedOrder := dto.NewOrderWithDetails(order, orderItems)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"order":   castedOrder,
	})
}
