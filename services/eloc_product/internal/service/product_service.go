package service

import (
	"context"
	"encoding/json"

	db "github.com/ryannguyen1105/eloc-backend/services/eloc_product/db/sqlc"
	"github.com/sqlc-dev/pqtype"
)

type CreateProductDTO struct {
	CategoryID int64
	Name       string
	Slug       string
	Sku        string
	Price      int64
	Stock      int32
	Attributes pqtype.NullRawMessage
}

func (productService *ProductService) CreateProduct(ctx context.Context, dto CreateProductDTO) (db.Product, error) {
	category, err := productService.CreateCategory(ctx, CreateCategoryDTO{
		Name: dto.Name,
		Slug: dto.Slug,
	})
	if err != nil {
		return db.Product{}, err
	}

	arg := db.CreateProductParams{
		CategoryID: category.ID,
		Name:       dto.Name,
		Slug:       dto.Slug,
		Sku:        dto.Sku,
		Price:      dto.Price,
		Stock:      dto.Stock,
		Attributes: pqtype.NullRawMessage{
			RawMessage: json.RawMessage([]byte("null")),
			Valid:      false,
		},
	}
	return productService.store.CreateProduct(ctx, arg)
}

type GetProductDTO struct {
	ID int64
}

func (productService *ProductService) GetProduct(ctx context.Context, dto GetProductDTO) (db.Product, error) {
	arg := db.GetProductByIDParams{
		ID: dto.ID,
	}
	return productService.store.GetProductByID(ctx, arg)
}

type UpdateProductDTO struct {
	CategoryID int64
	ID         int64
	Name       string
	Slug       string
	Sku        string
	Price      int64
	Stock      int32
}

func (productService *ProductService) UpdateProduct(ctx context.Context, dto UpdateProductDTO) (db.Product, error) {
	arg := db.UpdateProductParams{
		CategoryID: dto.CategoryID,
		ID:         dto.ID,
		Name:       dto.Name,
		Slug:       dto.Slug,
		Sku:        dto.Sku,
		Price:      dto.Price,
		Stock:      dto.Stock,
	}
	return productService.store.UpdateProduct(ctx, arg)
}

type UpdateProductStockDTO struct {
	ID    int64
	Stock int32
}

func (productService *ProductService) UpdateProductStock(ctx context.Context, dto UpdateProductStockDTO) (db.Product, error) {
	arg := db.UpdateProductStockParams{
		ID:    dto.ID,
		Stock: dto.Stock,
	}
	return productService.store.UpdateProductStock(ctx, arg)
}

type DeleteProductDTO struct {
	ID int64
}

func (productService *ProductService) DeleteProduct(ctx context.Context, dto DeleteProductDTO) (db.Product, error) {
	product, err := productService.store.GetProductByID(ctx, db.GetProductByIDParams{
		ID: dto.ID,
	})
	if err != nil {
		return db.Product{}, err
	}
	arg := db.DeleteProductParams{
		ID: product.ID,
	}
	return productService.store.DeleteProduct(ctx, arg)
}
