package db

import (
	"context"
	"testing"

	"github.com/ryannguyen1105/eloc-backend/util"
	"github.com/stretchr/testify/require"
)

func TestOrderTx(t *testing.T) {
	store := NewStore(testDB)

	userID := util.RandomUserID()
	productID1 := util.RandomProductID()
	productID2 := util.RandomProductID()

	_, err := store.CreateCart(context.Background(), CreateCartParams{
		UserID:    userID,
		ProductID: productID1,
		Quantity:  2,
	})
	require.NoError(t, err)

	_, err = store.CreateCart(context.Background(), CreateCartParams{
		UserID:    userID,
		ProductID: productID2,
		Quantity:  1,
	})
	require.NoError(t, err)

	n := 5
	errs := make(chan error, n)
	results := make(chan OrderTxResult, n)

	arg := OrderTxParams{
		UserID:          userID,
		TotalAmount:     util.RandomPrice(),
		Status:          "PENDING",
		ShippingAddress: util.RandomAddress(),
		CustomerPhone:   util.RandomPhone(),
		Items: []ItemOrderArg{
			{ProductID: productID1, Quantity: 2, Price: util.RandomPrice()},
			{ProductID: productID2, Quantity: 1, Price: util.RandomPrice()},
		},
	}
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.OrderTx(context.Background(), arg)
			errs <- err
			results <- result
		}()
	}
	successCount := 0
	for i := 0; i < n; i++ {
		err := <-errs
		result := <-results

		if err == nil {
			successCount++

			require.NotEmpty(t, result.Order)
			require.Equal(t, userID, result.Order.UserID)
			require.Equal(t, arg.TotalAmount, result.Order.TotalAmount)
			require.Equal(t, "PENDING", result.Order.Status)

			require.Len(t, result.OrderItem, 2)
			require.Equal(t, result.Order.ID, result.OrderItem[0].OrderID)
			require.Equal(t, result.Order.ID, result.OrderItem[1].OrderID)

		} else {
			require.NoError(t, err)
		}
	}

	require.Equal(t, n , successCount)

	cartItems, err := store.GetUserCart(context.Background(), GetUserCartParams{
		UserID: userID,
	})
	require.NoError(t, err)
	require.Empty(t, cartItems)
}
