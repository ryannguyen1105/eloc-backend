package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_order/db/sqlc"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_order/internal/service"
)

type createOrderRequest struct {
	UserID          int64  `json:"user_id" binding:"required,min=1"`
	ShippingAddress string `json:"shipping_address" binding:"required"`
	CustomerPhone   string `json:"customer_phone" binding:"required,min=1"`
}

type orderResponse struct {
	ID              int64     `json:"id"`
	UserID          int64     `json:"user_id"`
	TotalAmount     int64     `json:"total_amount"`
	Status          string    `json:"status"`
	ShippingAddress string    `json:"shipping_address"`
	CustomerPhone   string    `json:"customer_phone"`
	CreatedAt       time.Time `json:"created_at"`
}

func newOrderResponse(order db.Order) orderResponse {
	return orderResponse{
		ID:              order.ID,
		UserID:          order.UserID,
		TotalAmount:     order.TotalAmount,
		Status:          order.Status,
		ShippingAddress: order.ShippingAddress,
		CustomerPhone:   order.CustomerPhone,
		CreatedAt:       order.CreatedAt,
	}
}

func (server *Server) createOrder(ctx *gin.Context) {
	var req createOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.CreateOrderDTO{
		UserID:          req.UserID,
		ShippingAddress: req.ShippingAddress,
		CustomerPhone:   req.CustomerPhone,
	}
	order, err := server.orderService.CreateOrder(ctx, dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newOrderResponse(order)
	ctx.JSON(http.StatusOK, rsp)
}

type getOrderRequest struct {
	ID int64 `json:"id" binding:"required,min=1"`
}

func (server *Server) getOrder(ctx *gin.Context) {
	var req getOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.GetOrderDTO{
		ID: req.ID,
	}
	order, err := server.orderService.GetOrder(ctx, dto)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newOrderResponse(order)
	ctx.JSON(http.StatusOK, rsp)
}

type listOrderRequest struct {
	UserID int64 `json:"user_id" binding:"required,min=1"`
	Page   int32 `form:"page" binding:"required,min=1"`
	Limit  int32 `form:"limit" binding:"required,min=5,max=10"`
}

func (server *Server) listOrder(ctx *gin.Context) {
	var req listOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.ListOrderDTO{
		UserID: req.UserID,
		Page:   req.Page,
		Limit:  req.Limit,
	}
	order, err := server.orderService.ListOrder(ctx, dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var rsp []orderResponse = make([]orderResponse, 0)
	for _, orders := range order {
		convertedOrder := db.Order{
			ID:              orders.ID,
			UserID:          orders.UserID,
			TotalAmount:     orders.TotalAmount,
			Status:          orders.Status,
			ShippingAddress: orders.ShippingAddress,
			CustomerPhone:   orders.CustomerPhone,
			CreatedAt:       orders.CreatedAt,
		}

		rsp = append(rsp, newOrderResponse(convertedOrder))
	}

	ctx.JSON(http.StatusOK, rsp)
}

type updateOrderStatusRequest struct {
	ID     int64  `json:"id" binding:"required,min=1"`
	Status string `json:"status" binding:"required,oneof= pending shipping delivered"`
}

type updateOrderStatusResponse struct {
	ID              int64     `json:"id"`
	UserID          int64     `json:"user_id"`
	TotalAmount     int64     `json:"total_amount"`
	Status          string    `json:"status"`
	ShippingAddress string    `json:"shipping_address"`
	CustomerPhone   string    `json:"customer_phone"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func newUpdateOrderStatusResponse(order db.Order) updateOrderStatusResponse{
	return updateOrderStatusResponse{
		ID: order.ID,
		UserID: order.UserID,
		TotalAmount: order.TotalAmount,
		Status: order.Status,
		ShippingAddress: order.ShippingAddress,
		CustomerPhone: order.CustomerPhone,
		UpdatedAt: order.UpdatedAt,

	}
} 

func (server *Server) updateOrderStatus (ctx *gin.Context) {
	var req updateOrderStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.UpdateOrderStatusDTO{
		ID: req.ID,
		Status: req.Status,
	}
	order, err := server.orderService.UpdateOrderStatus(ctx, dto)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newUpdateOrderStatusResponse(order)
	ctx.JSON(http.StatusOK, rsp)
}