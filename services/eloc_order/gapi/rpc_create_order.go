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

func (server *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	cart, err := server.store.GetUserCart(ctx, db.GetUserCartParams{
		UserID: req.GetUserId(),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to find user")
	}
	if len(cart) == 0 {
		return nil, status.Errorf(codes.NotFound, "cart not found")
	}

	var totalAmount int64 = 0
	for _, item := range cart {
		var mockProductPrice int64 = 150000

		totalAmount += item.Quantity * mockProductPrice
	}

	arg := db.CreateOrderParams{
		UserID:          req.GetUserId(),
		TotalAmount:     totalAmount,
		Status:          "pending",
		ShippingAddress: req.GetShippingAddress(),
		CustomerPhone:   req.CustomerPhone,
	}
	order, err := server.store.CreateOrder(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed create order")
	}
	rsp := &pb.CreateOrderResponse{
		Order: convertOrder(order),
	}

	return rsp, nil
}
