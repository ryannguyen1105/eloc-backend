package gapi

import (
	pb "github.com/ryannguyen1105/eloc-backend/api/pb/eloc_order"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_order/db/sqlc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertCart(cart db.Cart) *pb.Cart {
	return &pb.Cart{
		UserId: cart.UserID,
		ProductId: cart.ProductID,
		Quantity: cart.Quantity,
	}
}

func convertOrder(order db.Order) *pb.Order {
	return &pb.Order{
		UserId: order.UserID,
		TotalAmount: order.TotalAmount,
		Status: order.Status,
		ShppingAddress: order.ShippingAddress,
		CustomerPhone: order.CustomerPhone,
		CreatedAt: timestamppb.New(order.CreatedAt),
	}
}