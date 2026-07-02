package service

import (
	"context"

	db "github.com/ryannguyen1105/eloc-backend/services/eloc_product/db/sqlc"
)

type ProductService struct {
	store db.Store
}

type CreateCategoryDTO struct {
	Name string
	Slug string
}

type GetCategoryDTO struct {
	Name string
}

type DeleteCategoryDTO struct {
	Name string
}

func NewProductService(store db.Store) *ProductService {
	return &ProductService{store: store}
}

func (productService *ProductService) CreateCategory(ctx context.Context, dto CreateCategoryDTO) (db.Category, error) {
	arg := db.CreateCategoryParams{
		Name: dto.Name,
		Slug: dto.Slug,
	}
	return productService.store.CreateCategory(ctx, arg)
}

func (productService *ProductService) GetCategory(ctx context.Context, dto GetCategoryDTO) (db.Category, error) {
	arg := db.GetCategoryByNameParams{
		Name: dto.Name,
	}
	return productService.store.GetCategoryByName(ctx, arg)
}

func (productService *ProductService) DeleteCategory(ctx context.Context, dto DeleteCategoryDTO) (db.Category, error) {
	category, err := productService.store.GetCategoryByName(ctx, db.GetCategoryByNameParams{
		Name: dto.Name,
	})
	if err != nil {
		return db.Category{}, err
	}
	err = productService.store.DeleteCategory(ctx, db.DeleteCategoryParams{
		Name: category.Name,
	})
	if err != nil {
		return db.Category{}, err
	}
	return category, nil
}