package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_order/db/sqlc"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_order/internal/service"
)

type cartResponse struct {
	UserID    int64     `json:"user_id" `
	ProductID int64     `json:"product_id" `
	Quantity  int64     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}

func newCartResponse(cart db.Cart) cartResponse {
	return cartResponse{
		UserID:    cart.UserID,
		ProductID: cart.ProductID,
		Quantity:  cart.Quantity,
		CreatedAt: cart.CreatedAt,
	}
}

type createCartRequest struct {
	UserID    int64 `json:"user_id" binding:"required,min=1"`
	ProductID int64 `json:"product_id" binding:"required,min=1"`
	Quantity  int64 `json:"quantity" binding:"required,min=1"`
}

func (server *Server) createCart(ctx *gin.Context) {
	var req createCartRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.CreateCartDTO{
		UserID:    req.UserID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}
	cart, err := server.orderService.CreateCart(ctx, dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newCartResponse(cart)
	ctx.JSON(http.StatusOK, rsp)
}

type getCartRequest struct {
	UserID int64 `json:"user_id" binding:"required,min=1"`
}

func (server *Server) getCart(ctx *gin.Context) {
	var req getCartRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.GetCartDTO{
		UserID: req.UserID,
	}
	cart, err := server.orderService.GetCart(ctx, dto)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, cart)
}

type updateCartQuantityResponse struct {
	UserID    int64     `json:"user_id"`
	ProductID int64     `json:"product_id"`
	Quantity  int64     `json:"quantity"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newUpdateCartQuantityResponse(cart db.Cart) updateCartQuantityResponse {
	return updateCartQuantityResponse{
		UserID:    cart.UserID,
		ProductID: cart.ProductID,
		Quantity:  cart.Quantity,
		UpdatedAt: cart.UpdatedAt,
	}
}

type updateCartQuantityRequest struct {
	UserID    int64 `json:"user_id" binding:"required,min=1"`
	ProductID int64 `json:"product_id" binding:"required,min=1"`
	Quantity  int64 `json:"quantity" binding:"required,min=1"`
}

func (server *Server) updateCartQuantity(ctx *gin.Context) {
	var req updateCartQuantityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.UpdateCartQuantityDTO{
		UserID:    req.UserID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}
	cart, err := server.orderService.UpdateCartQuantity(ctx, dto)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newUpdateCartQuantityResponse(cart)
	ctx.JSON(http.StatusOK, rsp)
}

type removeFromCartRequest struct {
	UserID    int64 `json:"user_id" binding:"required,min=1"`
	ProductID int64 `json:"product_id" binding:"required,min=1"`
}

func (server *Server) removeFromCart(ctx *gin.Context) {
	var req removeFromCartRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.RemoveFromCartDTO{
		UserID:    req.UserID,
		ProductID: req.ProductID,
	}
	cart, err := server.orderService.RemoveFromCart(ctx, dto)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "deleted successful", "data": cart})
}

type clearUserCartRequest struct {
	UserID int64 `json:"user_id" binding:"required,min=1"`
}

func (server *Server) clearFromCart(ctx *gin.Context) {
	var req clearUserCartRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.ClearUserCartDTO{
		UserID: req.UserID,
	}
	cart, err := server.orderService.ClearUserCart(ctx, dto)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "deleted successful", "data": cart})
}
