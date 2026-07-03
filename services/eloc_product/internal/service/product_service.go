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

type GetProductDTO struct {
	Name string
}

type UpdateProductDTO struct {
	ID         int64
	CategoryID int64
	Name       string
	Slug       string
	Sku        string
	Price      int64
	Stock      int32
}

type UpdateProductStockDTO struct {
	Name  string
	Stock int32
}

type DeleteProductDTO struct {
	Name string
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

func (productService *ProductService) GetProduct(ctx context.Context, dto GetProductDTO) (db.Product, error) {
	arg := db.GetProductByNameParams{
		Name: dto.Name,
	}
	return productService.store.GetProductByName(ctx, arg)
}

func (productService *ProductService) UpdateProduct(ctx context.Context, dto UpdateProductDTO) (db.Product, error) {
	arg := db.UpdateProductParams{
		ID: dto.ID,
		CategoryID: dto.ID,
		Name: dto.Name,
		Slug: dto.Slug,
		Sku: dto.Sku,
		Price: dto.Price,
		Stock: dto.Stock,
	}
	return productService.store.UpdateProduct(ctx, arg)
}

func (productService *ProductService) UpdateProductStock(ctx context.Context, dto UpdateProductStockDTO) (db.Product, error) {
	arg := db.UpdateProductStockParams{
		Name: dto.Name,
		Stock: dto.Stock,
	}
	return productService.store.UpdateProductStock(ctx, arg)
}

func (productService *ProductService) DeleteProduct(ctx context.Context, dto DeleteProductDTO) (db.Product, error) {
	product, err := productService.store.GetProductByName(ctx, db.GetProductByNameParams{
		Name: dto.Name,
	})
	if err != nil {
		return db.Product{}, err
	}
	err = productService.store.DeleteProduct(ctx, db.DeleteProductParams{
		Name: dto.Name,
	})
	if err != nil {
		return db.Product{}, err
	}
	return product, nil
}
