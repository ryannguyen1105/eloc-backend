package gapi

import (
	"context"
	"database/sql"
	"errors"

	pb "github.com/ryannguyen1105/eloc-backend/api/pb/eloc_order"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_order/db/sqlc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateCartQuantity(ctx context.Context, req *pb.UpdateCartQuantityRequest) (*pb.UpdateCartQuantityResponse, error) {
	arg := db.UpdateCartQuantityParams{
		UserID:    req.GetUserId(),
		ProductID: req.GetProductId(),
		Quantity:  req.GetQuantity(),
	}
	cart, err := server.store.UpdateCartQuantity(ctx, arg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to find user")
	}
	rsp := &pb.UpdateCartQuantityResponse{
		Cart: convertCart(cart),
	}

	return rsp, nil
}
