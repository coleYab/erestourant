package store

import (
	"context"
	"fmt"

	"github.com/coleYab/erestourant/internal/db/repository"
	"github.com/coleYab/erestourant/internal/dto"
	"github.com/google/uuid"
)

type OrderStore struct {
	qry *repository.Queries
}

func NewOrderStore(qry *repository.Queries) *OrderStore {
	return &OrderStore{qry: qry}
}

func (s *OrderStore) VerifyOrder(orderId string) error {
	uid, _ := uuid.Parse(orderId)
	ctx := context.Background()
	return s.qry.VerfiyOrder(ctx, uid)
}

func (s *OrderStore) CreateOrder(orderParams dto.CreateOrderDto, tot float64, id string) (repository.Order, error) {
	uid, _ := uuid.Parse(id)
	ctx := context.Background()

	order, err := s.qry.CreateOrder(ctx, repository.CreateOrderParams{
		UserId:     uid,
		TotalPrice: tot,
		Status:     "in progress",
	})

	if err != nil {
		return repository.Order{}, fmt.Errorf("unable to create a new order %v", err.Error())
	}

	for _, mitem := range orderParams.Items {
		uid, _ := uuid.Parse(mitem.Id)
		menuItem, _ := s.qry.GetMenuById(ctx, uid)
		s.qry.UpdateMenuQtyById(ctx, repository.UpdateMenuQtyByIdParams{ID: menuItem.ID, Qty: menuItem.Qty - int32(mitem.Qty)})
		_ = s.qry.CreateOrderItem(ctx, repository.CreateOrderItemParams{
			OrderId:    order.ID,
			MenuItemId: menuItem.ID,
			UnitPrice:  menuItem.Price,
			Qty:        menuItem.Qty,
		})
	}

	return order, nil
}

func (s *OrderStore) GetOrderByID(id string) (repository.Order, error) {
	uid, _ := uuid.Parse(id)
	ctx := context.Background()
	return s.qry.GetOrderById(ctx, uid)
}

func (s *OrderStore) GetOrderItemsByOrderId(id string) ([]repository.OrderItem, error) {
	uid, _ := uuid.Parse(id)
	ctx := context.Background()
	return s.qry.GetOrderItems(ctx, uid)
}

func (s *OrderStore) GetAllOrder() ([]repository.Order, error) {
	ctx := context.Background()
	return s.qry.GetAllOrders(ctx)
}

func (s *OrderStore) DeleteOrder(id string) error {
	uid, _ := uuid.Parse(id)
	ctx := context.Background()
	return s.qry.DeleteMenuById(ctx, uid)
}
