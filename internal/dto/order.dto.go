package dto

import (
	"github.com/coleYab/erestourant/internal/db/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateOrderDto struct {
	Items []OrderItem `json:"items" binding:"required,dive"`
}

type OrderItem struct {
	Id  string `json:"id" binding:"required,uuid"`
	Qty int    `json:"qty" binding:"required,gt=0"`
}

type OrderWithDeatils struct {
	ID         uuid.UUID        `json:"id"`
	TotalPrice float64          `json:"totalPrice"`
	UserId     uuid.UUID        `json:"userId"`
	Status     string           `json:"status"`
	CreatedAt  pgtype.Timestamp `json:"createdAt"`
	UpdatedAt  pgtype.Timestamp `json:"updatedAt"`
	OrderItems []repository.OrderItem `json:"orderItems"`
}

func NewOrderWithDetails(order repository.Order, orderItems []repository.OrderItem) OrderWithDeatils {
	return OrderWithDeatils{
		ID: order.ID,
		TotalPrice: order.TotalPrice,
		UserId: order.UserId,
		Status: order.Status,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
		OrderItems: orderItems,
	}
}
