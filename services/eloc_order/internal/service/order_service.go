package service

import (
	"context"
	"errors"
	"time"

	db "github.com/ryannguyen1105/eloc-backend/services/eloc_order/db/sqlc"
)

type CreateOrderDTO struct {
	UserID          int64
	Status          string
	ShippingAddress string
	CustomerPhone   string
}

type GetOrderDTO struct {
	ID int64
}

type ListOrderDTO struct {
	UserID int64
	Page   int32
	Limit  int32
}

type UpdateOrderStatusDTO struct {
	ID        int64
	Status    string
	UpdatedAt time.Time
}

func (orderService *OrderService) CreateOrder(ctx context.Context, dto CreateOrderDTO) (db.Order, error) {
	cart, err := orderService.store.GetUserCart(ctx, db.GetUserCartParams{
		UserID: dto.UserID,
	})
	if err != nil {
		return db.Order{}, err
	}
	if len(cart) == 0 {
		return db.Order{}, errors.New("can't create order")
	}

	var totalAmount int64 = 0
	for _, item := range cart {
		var mockProductPrice int64 = 150000

		totalAmount += item.Quantity * mockProductPrice
	}

	arg := db.CreateOrderParams{
		UserID:          dto.UserID,
		TotalAmount:     totalAmount,
		Status:          "pending",
		ShippingAddress: dto.ShippingAddress,
		CustomerPhone:   dto.CustomerPhone,
	}
	return orderService.store.CreateOrder(ctx, arg)
}

func (orderService *OrderService) GetOrder(ctx context.Context, dto GetOrderDTO) (db.Order, error) {
	arg := db.GetOrderParams{
		ID: dto.ID,
	}
	return orderService.store.GetOrder(ctx, arg)
}

func (orderService *OrderService) ListOrder(ctx context.Context, dto ListOrderDTO) ([]db.ListUserOrdersRow, error) {
	offset := (dto.Page - 1) * dto.Limit

	arg := db.ListUserOrdersParams{
		UserID: dto.UserID,
		Limit:  dto.Limit,
		Offset: offset,
	}
	return orderService.store.ListUserOrders(ctx, arg)
}

func (orderService *OrderService) UpdateOrderStatus (ctx context.Context, dto UpdateOrderStatusDTO) (db.Order, error) {
	arg := db.UpdateOrderStatusParams{
		ID: dto.ID,
		Status: dto.Status,
	}
	return orderService.store.UpdateOrderStatus(ctx, arg)
}