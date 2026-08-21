package service

import (
	"context"

	db "github.com/ryannguyen1105/eloc-backend/services/eloc_order/db/sqlc"
)

type CreateCartDTO struct {
	UserID    int64
	ProductID int64
	Quantity  int64
}

func (orderService *OrderService) CreateCart(ctx context.Context, dto CreateCartDTO) (db.Cart, error) {
	arg := db.CreateCartParams{
		UserID:    dto.UserID,
		ProductID: dto.ProductID,
		Quantity:  dto.Quantity,
	}
	return orderService.store.CreateCart(ctx, arg)
}

type GetCartDTO struct {
	UserID int64
}

func (orderService *OrderService) GetCart(ctx context.Context, dto GetCartDTO) ([]db.Cart, error) {
	arg := db.GetUserCartParams{
		UserID: dto.UserID,
	}
	return orderService.store.GetUserCart(ctx, arg)
}

type UpdateCartQuantityDTO struct {
	UserID    int64
	ProductID int64
	Quantity  int64
}

func (orderService *OrderService) UpdateCartQuantity(ctx context.Context, dto UpdateCartQuantityDTO) (db.Cart, error) {
	arg := db.UpdateCartQuantityParams{
		UserID:    dto.UserID,
		ProductID: dto.ProductID,
		Quantity:  dto.Quantity,
	}
	return orderService.store.UpdateCartQuantity(ctx, arg)
}

type RemoveFromCartDTO struct {
	UserID    int64
	ProductID int64
}

func (orderService *OrderService) RemoveFromCart(ctx context.Context, dto RemoveFromCartDTO) (db.Cart, error) {
	_, err := orderService.store.GetUserCart(ctx, db.GetUserCartParams{
		UserID: dto.UserID,
	})
	if err != nil {
		return db.Cart{}, err
	}
	arg := db.RemoveFromCartParams{
		UserID:    dto.UserID,
		ProductID: dto.ProductID,
	}
	return orderService.store.RemoveFromCart(ctx, arg)
}

type ClearUserCartDTO struct {
	UserID int64
}

func (orderService *OrderService) ClearUserCart(ctx context.Context, dto ClearUserCartDTO) (db.Cart, error) {
	_, err := orderService.store.ClearUserCart(ctx, db.ClearUserCartParams{
		UserID: dto.UserID,
	})
	if err != nil {
		return db.Cart{}, err
	}
	arg := db.ClearUserCartParams{
		UserID: dto.UserID,
	}
	return orderService.store.ClearUserCart(ctx, arg)
}
