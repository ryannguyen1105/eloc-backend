package service

import (
	"context"

	db "github.com/ryannguyen1105/eloc-backend/services/eloc_product/db/sqlc"
)

type AddProductImageDTO struct {
	ProductID int64
	ImageUrl  string
	IsPrimary bool
}

func (productService *ProductService) AddProductImage(ctx context.Context, dto AddProductImageDTO) (db.ProductImage, error) {
	arg := db.AddProductImageParams{
		ProductID: dto.ProductID,
		ImageUrl:  dto.ImageUrl,
		IsPrimary: dto.IsPrimary,
	}
	return productService.store.AddProductImage(ctx, arg)
}

type DeleteProductImageDTO struct {
	ID        int64
	ProductID int64
}

func (productService *ProductService) DeleteProductImage(ctx context.Context, dto DeleteProductImageDTO) (db.ProductImage, error) {
	arg := db.DeleteProductImageParams{
		ID:        dto.ID,
		ProductID: dto.ProductID,
	}
	return db.ProductImage{}, productService.store.DeleteProductImage(ctx, arg)
}
