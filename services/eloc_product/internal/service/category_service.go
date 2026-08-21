package service

import (
	"context"

	db "github.com/ryannguyen1105/eloc-backend/services/eloc_product/db/sqlc"
)

type CreateCategoryDTO struct {
	Name string
	Slug string
}

func (productService *ProductService) CreateCategory(ctx context.Context, dto CreateCategoryDTO) (db.Category, error) {
	arg := db.CreateCategoryParams{
		Name: dto.Name,
		Slug: dto.Slug,
	}
	return productService.store.CreateCategory(ctx, arg)
}

type GetCategoryDTO struct {
	ID int64
}

func (productService *ProductService) GetCategory(ctx context.Context, dto GetCategoryDTO) (db.Category, error) {
	arg := db.GetCategoryByIDParams{
		ID: dto.ID,
	}
	return productService.store.GetCategoryByID(ctx, arg)
}

type DeleteCategoryDTO struct {
	ID int64
}

func (productService *ProductService) DeleteCategory(ctx context.Context, dto DeleteCategoryDTO) (db.Category, error) {
	category, err := productService.store.GetCategoryByID(ctx, db.GetCategoryByIDParams{
		ID: dto.ID,
	})
	if err != nil {
		return db.Category{}, err
	}
	err = productService.store.DeleteCategory(ctx, db.DeleteCategoryParams{
		ID: category.ID,
	})
	if err != nil {
		return db.Category{}, err
	}
	return category, nil
}
