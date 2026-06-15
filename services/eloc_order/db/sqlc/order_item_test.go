package db

import (
	"context"
	"testing"

	"github.com/ryannguyen1105/eloc-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomOrderItem(t *testing.T, cart Cart, order Order) OrderItem {
	arg := CreateOrderItemParams{
		OrderID:   order.ID,
		ProductID: cart.ProductID,
		Quantity:  cart.Quantity,
		Price:     util.RandomPrice(),
	}
	orderItem, err := testQueries.CreateOrderItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orderItem)

	return orderItem
}

func TestCreateOrderItem(t *testing.T) {
	cart := createRandomCarts(t)
	order := createRandomOrder(t, cart)
	createRandomOrderItem(t, cart, order)
}

func TestGetOrderItems(t *testing.T) {
	cart := createRandomCarts(t)
	order := createRandomOrder(t, cart)
	orderItem1 := createRandomOrderItem(t, cart, order)

	arg := GetOrderItemsParams {
		OrderID: orderItem1.OrderID,
	}
	orderItem2, err := testQueries.GetOrderItems(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orderItem2)
}
