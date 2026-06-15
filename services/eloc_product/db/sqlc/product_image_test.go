package db

import (
	"context"
	"testing"

	"github.com/ryannguyen1105/eloc-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomProductImage(t *testing.T, product Product, isPrimary bool) ProductImage {
	arg := AddProductImageParams{
		ProductID: product.ID,
		ImageUrl:  util.RandomUrl(),
		IsPrimary: isPrimary,
	}
	image, err := testQueries.AddProductImage(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, image)

	require.Equal(t, arg.ProductID, image.ProductID)
	require.Equal(t, arg.ImageUrl, image.ImageUrl)
	require.Equal(t, arg.IsPrimary, image.IsPrimary)

	return image
}

func TestAddProductImage(t *testing.T) {
	category := createRandomCategory(t)
	product := createRandomProduct(t, category)
	createRandomProductImage(t, product, true)
}

func TestGetProductImage(t *testing.T) {
	category := createRandomCategory(t)
	product := createRandomProduct(t, category)
	image1 := createRandomProductImage(t, product, true)

	arg := GetProductImagesParams{
		ProductID: image1.ProductID,
	}

	images, err := testQueries.GetProductImages(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, images)

	image2 := images[0]

	require.Equal(t, image1.ID, image2.ID)
	require.Equal(t, image1.ProductID, image2.ProductID)
	require.Equal(t, image1.ImageUrl, image2.ImageUrl)
	require.Equal(t, image1.IsPrimary, image2.IsPrimary)

}

func TestResetPrimaryImage(t *testing.T) {
	category := createRandomCategory(t)
	product := createRandomProduct(t, category)
	createRandomProductImage(t, product, true)

	arg := ResetPrimaryImageParams{
		ProductID: product.ID,
	}

	err := testQueries.ResetPrimaryImage(context.Background(), arg)
	require.NoError(t, err)

	getArg := GetProductImagesParams{
		ProductID: product.ID,
	}
	images, err := testQueries.GetProductImages(context.Background(), getArg)
	require.NoError(t, err)
	require.NotEmpty(t, images)

	for _, img := range images {
		require.False(t, img.IsPrimary)
	}
}

func TestSetPrimaryImage(t *testing.T) {
	category := createRandomCategory(t)
	product := createRandomProduct(t, category)
	image1 := createRandomProductImage(t, product, false)

	arg := SetPrimaryImageParams{
		ID:        image1.ID,
		ProductID: product.ID,
	}

	updateImg, err := testQueries.SetPrimaryImage(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updateImg)
	require.True(t, updateImg.IsPrimary)

	getArg := GetProductImagesParams{
		ProductID: product.ID,
	}
	image2, err := testQueries.GetProductImages(context.Background(), getArg)
	require.NoError(t, err)
	require.NotEmpty(t, image2)

	for _, image := range image2 {
		require.True(t, image.IsPrimary)
	}
}

func TestDeleteProductImage(t *testing.T) {
	category := createRandomCategory(t)
	product := createRandomProduct(t, category)
	image1 := createRandomProductImage(t, product, false)

	deleteArg := DeleteProductImageParams{
		ID:        image1.ID,
		ProductID: product.ID,
	}
	err := testQueries.DeleteProductImage(context.Background(), deleteArg)
	require.NoError(t, err)

	arg := GetProductImagesParams{
		ProductID: image1.ProductID,
	}
	image2, err := testQueries.GetProductImages(context.Background(), arg)
	require.NoError(t, err)

	require.Empty(t, image2)
	require.Len(t, image2, 0)

}
