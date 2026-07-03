package service

import db "github.com/ryannguyen1105/eloc-backend/services/eloc_product/db/sqlc"

type ProductService struct {
	store db.Store
}

func NewProductService(store db.Store) *ProductService {
	return &ProductService{store: store}
}
