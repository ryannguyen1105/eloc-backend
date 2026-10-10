package gapi

import (
	pb "github.com/ryannguyen1105/eloc-backend/api/pb/eloc_product"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_product/db/sqlc"
)

func convertProduct(product db.Product) *pb.Product {
	return &pb.Product{
		Name: product.Name,
		Price: product.Price,
		Stock: product.Stock,
		Description: product.Description,
		Status: product.Status,
	}
}