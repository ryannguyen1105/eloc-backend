package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(q *Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rollback: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type ItemOrderArg struct {
	ProductID int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
	Price     int64 `json:"price"`
}

type OrderTxParams struct {
	UserID          int64          `json:"user_id"`
	TotalAmount     int64          `json:"total_amount"`
	Status          string         `json:"status"`
	ShippingAddress string         `json:"shipping_address"`
	CustomerPhone   string         `json:"customer_phone"`
	Items           []ItemOrderArg `json:"items"`
}

type OrderTxResult struct {
	Order     Order       `json:"order"`
	OrderItem []OrderItem `json:"order_items"`
}

func (store *Store) OrderTx(ctx context.Context, arg OrderTxParams) (OrderTxResult, error) {
	var result OrderTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Order, err = q.CreateOrder(ctx, CreateOrderParams{
			UserID: arg.UserID,
			TotalAmount: arg.TotalAmount,
			Status: arg.Status,
			ShippingAddress: arg.ShippingAddress,
			CustomerPhone: arg.CustomerPhone,
		})
		if err != nil {
			return err
		}
		result.OrderItem = make([]OrderItem, len(arg.Items))
		for i, item := range arg.Items {
			result.OrderItem[i], err = q.CreateOrderItem(ctx, CreateOrderItemParams{
				OrderID: result.Order.ID,
				ProductID: item.ProductID,
				Quantity: item.Quantity,
				Price: item.Quantity,
			})
			if err != nil {
				return err
			}
		}
		err = q.ClearUserCart(ctx, ClearUserCartParams{
			UserID: arg.UserID,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}