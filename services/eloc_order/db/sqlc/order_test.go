package db

import (
	"context"
	"testing"
	"time"

	"github.com/ryannguyen1105/eloc-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomOrder(t *testing.T, cart Cart) Order {
	arg := CreateOrderParams{
		UserID:          cart.UserID,
		TotalAmount:     util.RandomPrice(),
		Status:          util.RandomStatus(),
		ShippingAddress: util.RandomAddress(),
		CustomerPhone:   util.RandomPhone(),
	}
	order, err := testQueries.CreateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, arg.UserID, order.UserID)
	require.Equal(t, arg.TotalAmount, order.TotalAmount)
	require.Equal(t, arg.Status, order.Status)
	require.Equal(t, arg.ShippingAddress, order.ShippingAddress)
	require.Equal(t, arg.CustomerPhone, order.CustomerPhone)

	require.WithinDuration(t, time.Now(), order.CreatedAt, time.Second)

	return order

}

func TestCreateOrder(t *testing.T) {
	cart := createRandomCarts(t)
	createRandomOrder(t, cart)
}

func TestGetOrder(t *testing.T) {
	cart := createRandomCarts(t)
	order1 := createRandomOrder(t, cart)

	arg := GetOrderParams{
		ID: order1.ID,
	}
	order2, err := testQueries.GetOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order2)

	require.Equal(t, order1.ID, order2.ID)
	require.Equal(t, order1.UserID, order2.UserID)
	require.Equal(t, order1.TotalAmount, order2.TotalAmount)
	require.Equal(t, order1.Status, order2.Status)
	require.Equal(t, order1.ShippingAddress, order2.ShippingAddress)
	require.Equal(t, order1.CustomerPhone, order2.CustomerPhone)

	require.WithinDuration(t, order1.CreatedAt, order2.CreatedAt, time.Second)
	require.WithinDuration(t, order1.UpdatedAt, order2.UpdatedAt, time.Second)
}

func TestUpdateOrderStatus(t *testing.T) {
	cart := createRandomCarts(t)
	order1 := createRandomOrder(t, cart)

	arg := UpdateOrderStatusParams{
		ID:     order1.ID,
		Status: util.RandomStatus(),
	}
	order2, err := testQueries.UpdateOrderStatus(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order2)

	require.Equal(t, order1.ID, order2.ID)
	require.Equal(t, arg.Status, order2.Status)
	require.WithinDuration(t, time.Now(), order2.UpdatedAt, time.Second)
}

func TestListUserOrders(t *testing.T) {
	cart := createRandomCarts(t)
	for i := 0; i < 5; i++ {
		createRandomOrder(t, cart)
	}
	arg := ListUserOrdersParams{
		UserID: cart.UserID,
		Limit:  5,
		Offset: 0,
	}
	orders, err := testQueries.ListUserOrders(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, orders, 5)

	for _, order := range orders {
		require.NotEmpty(t, order)
		require.Equal(t, cart.UserID, order.UserID)
	}
}
