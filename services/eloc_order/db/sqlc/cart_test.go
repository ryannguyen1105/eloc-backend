package db

import (
	"context"
	"testing"
	"time"

	"github.com/ryannguyen1105/eloc-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomCarts(t *testing.T) Cart {
	arg := CreateCartParams{
		UserID:    util.RandomUserID(),
		ProductID: util.RandomProductID(),
		Quantity:  1,
	}
	cart, err := testQueries.CreateCart(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cart)

	require.Equal(t, arg.UserID, cart.UserID)
	require.Equal(t, arg.ProductID, cart.ProductID)
	require.Equal(t, arg.Quantity, cart.Quantity)

	return cart
}

func TestCreateCart(t *testing.T) {
	createRandomCarts(t)
}

func TestGetUserCart(t *testing.T) {
	cart1 := createRandomCarts(t)
	arg := GetUserCartParams{
		UserID: cart1.UserID,
	}
	cart2, err := testQueries.GetUserCart(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cart2)

	found := false
	for _, cart := range cart2 {
		if cart.ProductID == cart1.ProductID {
			require.Equal(t, cart1.UserID, cart.UserID)
			require.Equal(t, cart1.Quantity, cart.Quantity)
			found = true
			break
		}
	}
	require.True(t, found)
}

func TestUpdateCartQuantity(t *testing.T) {
	cart1 := createRandomCarts(t)
	arg := UpdateCartQuantityParams{
		UserID:    cart1.UserID,
		ProductID: cart1.ProductID,
		Quantity:  cart1.Quantity,
	}
	cart2, err := testQueries.UpdateCartQuantity(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cart2)

	require.Equal(t, arg.UserID, cart2.UserID)
	require.Equal(t, arg.ProductID, cart2.ProductID)
	require.Equal(t, arg.Quantity, cart2.Quantity)

	require.WithinDuration(t, time.Now(), cart2.UpdatedAt, time.Second)

}

func TestRemoveFromCart(t *testing.T) {
	cart1 := createRandomCarts(t)
	arg := RemoveFromCartParams{
		UserID:    cart1.UserID,
		ProductID: cart1.ProductID,
	}
	err := testQueries.RemoveFromCart(context.Background(), arg)
	require.NoError(t, err)

	getArg := GetUserCartParams{
		UserID: cart1.UserID,
	}
	cart2, err := testQueries.GetUserCart(context.Background(), getArg)
	require.NoError(t, err)

	for _, cart := range cart2 {
	require.NotEqual(t, cart1.ProductID, cart.ProductID)
	}
}

func TestClearUserCart(t *testing.T) {
	cart1 := createRandomCarts(t)
	arg := ClearUserCartParams{
		UserID: cart1.UserID,
	}
	err := testQueries.ClearUserCart(context.Background(), arg)
	require.NoError(t, err)

	getArg := GetUserCartParams{
		UserID: cart1.UserID,
	}
	cart2, err := testQueries.GetUserCart(context.Background(), getArg)
	require.NoError(t, err)

	require.Empty(t, cart2)
}
