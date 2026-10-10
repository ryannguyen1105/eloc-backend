package gapi

import (
	"context"
	"database/sql"
	"errors"

	pb "github.com/ryannguyen1105/eloc-backend/api/pb/eloc_product"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_product/db/sqlc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	arg := db.GetCategoryByIDParams{
		ID: req.GetId(),
	}
	product, err := server.store.GetProductByID(ctx, db.GetProductByIDParams(arg))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "product not foud")
		}

		return nil, status.Error(codes.Internal, "failed to find product")
	}
	rsp := &pb.GetProductResponse{
		Product: convertProduct(product),
	}
	return rsp, nil
}
