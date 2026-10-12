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

func (server *Server) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	arg := db.GetUserCartParams{
		UserID: req.GetUserId(),
	}
	cart, err := server.store.GetUserCart(ctx, arg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "faile to find user")
	}
	if len(cart) == 0 {
    return nil, status.Error(codes.NotFound, "cart not found")
}

return &pb.GetCartResponse{
    Cart: convertCart(cart[0]),
}, nil
}
