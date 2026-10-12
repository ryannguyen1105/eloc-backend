package gapi

import (
	"context"

	pb "github.com/ryannguyen1105/eloc-backend/api/pb/eloc_order"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_order/db/sqlc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateCart(ctx context.Context, req *pb.CreateCartRequest) (*pb.CreateCartResponse, error) {
	arg := db.CreateCartParams{
		UserID:    req.GetUserId(),
		ProductID: req.GetProductId(),
		Quantity:  req.GetQuantity(),
	}
	cart, err := server.store.CreateCart(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "faile to create cart")
	}
	rsp := &pb.CreateCartResponse{
		Cart: convertCart(cart),
	}
	return rsp, nil
}
