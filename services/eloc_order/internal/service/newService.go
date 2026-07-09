package service

import db "github.com/ryannguyen1105/eloc-backend/services/eloc_order/db/sqlc"

type OrderService struct {
	store db.Store
}

func NewOrderService(store db.Store) *OrderService {
	return &OrderService{store: store}
}