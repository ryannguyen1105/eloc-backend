package gapi

import (
	pb "github.com/ryannguyen1105/eloc-backend/api/pb/eloc_order"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_order/db/sqlc"
)

func convertCart(cart db.Cart) *pb.Cart {
	return &pb.Cart{
		UserId: cart.UserID,
		ProductId: cart.ProductID,
		Quantity: cart.Quantity,
	}
}