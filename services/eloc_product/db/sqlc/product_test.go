package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ryannguyen1105/eloc-backend/util"
	"github.com/sqlc-dev/pqtype"
	"github.com/stretchr/testify/require"
)

func createRandomProduct(t *testing.T, category Category) Product {
	productSlug := category.Slug + "-" + util.RandomSlug()
	arg := CreateProductParams{
		CategoryID: category.ID,
		Name:       category.Name,
		Slug:       productSlug,
		Sku:        util.RandomSku(),
		Price:      util.RandomPrice(),
		Stock:      0,
		Attributes: pqtype.NullRawMessage{
			RawMessage: []byte("null"),
			Valid:      false,
		},
	}
	product, err := testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.CategoryID, product.CategoryID)
	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Slug, product.Slug)
	require.Equal(t, arg.Sku, product.Sku)
	require.Equal(t, arg.Price, product.Price)
	require.Equal(t, arg.Stock, product.Stock)
	require.Equal(t, arg.Attributes.Valid, product.Attributes.Valid)

	return product
}

func TestCreateProduct(t *testing.T) {
	category := createRandomCategory(t)
	createRandomProduct(t, category)
}

func TestGetProduct(t *testing.T) {
	category := createRandomCategory(t)
	product1 := createRandomProduct(t, category)

	arg := GetProductByIDParams{
		ID: product1.ID,
	}
	product2, err := testQueries.GetProductByID(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.ID, product2.ID)
	require.Equal(t, product1.CategoryID, product2.CategoryID)
	require.Equal(t, product1.Name, product2.Name)
	require.Equal(t, product1.Slug, product2.Slug)
	require.Equal(t, product1.Slug, product2.Slug)
	require.Equal(t, product1.Sku, product2.Sku)
	require.Equal(t, product1.Price, product2.Price)
	require.Equal(t, product1.Stock, product2.Stock)
	require.Equal(t, product1.Attributes.Valid, product2.Attributes.Valid)
}

func TestGetProductBySlug(t *testing.T) {
	category := createRandomCategory(t)
	product1 := createRandomProduct(t, category)

	arg := GetProductBySlugParams{
		Slug: product1.Slug,
	}
	product2, err := testQueries.GetProductBySlug(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.ID, product2.ID)
	require.Equal(t, product1.Name, product2.Name)
	require.Equal(t, product1.Slug, product2.Slug)
	require.Equal(t, product1.Slug, product2.Slug)
	require.Equal(t, product1.Sku, product2.Sku)
	require.Equal(t, product1.Price, product2.Price)
	require.Equal(t, product1.Stock, product2.Stock)
	require.Equal(t, product1.Attributes.Valid, product2.Attributes.Valid)
}

func TestUpdateProduct(t *testing.T) {
	category := createRandomCategory(t)
	product1 := createRandomProduct(t, category)
	arg := UpdateProductParams{
		ID:         product1.ID,
		CategoryID: product1.CategoryID,
		Name:       product1.Name + " ",
		Slug:       product1.Slug,
		Sku:        product1.Sku,
		Price:      product1.Price,
		Stock:      product1.Stock,
		Attributes: product1.Attributes,
	}
	product2, err := testQueries.UpdateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, arg.ID, product2.ID)
	require.Equal(t, arg.Name, product2.Name)
	require.Equal(t, arg.Slug, product2.Slug)
	require.Equal(t, arg.Slug, product2.Slug)
	require.Equal(t, arg.Sku, product2.Sku)
	require.Equal(t, arg.Price, product2.Price)
	require.Equal(t, arg.Stock, product2.Stock)
	require.Equal(t, arg.Attributes.Valid, product2.Attributes.Valid)
}

func TestUpdateProductStock(t *testing.T) {
	category := createRandomCategory(t)
	product1 := createRandomProduct(t, category)
	arg := UpdateProductStockParams{
		ID:    product1.ID,
		Stock: product1.Stock,
	}
	product2, err := testQueries.UpdateProductStock(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, arg.ID, product2.ID)
	require.Equal(t, arg.Stock, product2.Stock)
}

func TestDeleteProduct(t *testing.T) {
	category := createRandomCategory(t)
	product1 := createRandomProduct(t, category)
	deleteArg := DeleteProductParams{
		ID: product1.ID,
	}
	err := testQueries.DeleteProduct(context.Background(), deleteArg)
	require.NoError(t, err)

	arg := GetProductByIDParams{
		ID: product1.ID,
	}
	product2, err := testQueries.GetProductByID(context.Background(), arg)
	require.Error(t, err)
	require.ErrorIs(t, err, sql.ErrNoRows)

	require.Empty(t, product2)
}

func TestListProducts(t *testing.T) {
	for i := 0; i < 5; i++ {
		category := createRandomCategory(t)
		createRandomProduct(t, category)

	}
	arg := ListProductsParams{
		Limit:  5,
		Offset: 5,
	}
	products, err := testQueries.ListProducts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, products, 5)

	for _, product := range products {
		require.NotEmpty(t, product)
	}
}

func TestListProductsByCategory(t *testing.T) {
	category := createRandomCategory(t)
	for i := 0; i < 5; i++ {
		createRandomProduct(t, category)

	}
	arg := ListProductsByCategoryParams{
		CategoryID: category.ID,
		Limit:      5,
		Offset:     0,
	}
	products, err := testQueries.ListProductsByCategory(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, products, 5)

	for _, product := range products {
		require.NotEmpty(t, product)
		require.Equal(t, category.ID, product.CategoryID)
	}
}
