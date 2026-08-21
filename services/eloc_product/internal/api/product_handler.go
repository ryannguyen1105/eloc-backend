package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_product/db/sqlc"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_product/internal/service"
	"github.com/sqlc-dev/pqtype"
)

type productResponse struct {
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Sku       string    `json:"sku"`
	Price     int64     `json:"price"`
	Stock     int32     `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
}

func newProductResponse(product db.Product) productResponse {
	return productResponse{
		Name:      product.Name,
		Slug:      product.Slug,
		Sku:       product.Sku,
		Price:     product.Price,
		Stock:     product.Stock,
		CreatedAt: product.CreatedAt,
	}
}

type createProductRequest struct {
	Name  string `json:"name" binding:"required"`
	Slug  string `json:"slug" binding:"required"`
	Sku   string `json:"sku" binding:"required"`
	Price int64  `json:"price" binding:"required,min=1"`
	Stock int32  `json:"stock" binding:"required,min=1"`
}

func (server *Server) createProduct(ctx *gin.Context) {
	var req createProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.CreateProductDTO{
		Name:  req.Name,
		Slug:  req.Slug,
		Sku:   req.Sku,
		Price: req.Price,
		Stock: req.Stock,
		Attributes: pqtype.NullRawMessage{
			RawMessage: json.RawMessage([]byte("null")),
			Valid:      false,
		},
	}
	product, err := server.productService.CreateProduct(ctx, dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newProductResponse(product)
	ctx.JSON(http.StatusOK, rsp)
}

type getProductRequest struct {
	ID int64 `json:"id" binding:"required,min=1"`
}

func (server *Server) getProduct(ctx *gin.Context) {
	var req getProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.GetProductDTO{
		ID: req.ID,
	}
	product, err := server.productService.GetProduct(ctx, dto)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newProductResponse(product)
	ctx.JSON(http.StatusOK, rsp)
}

type updateProductResponse struct {
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Sku       string    `json:"sku"`
	Price     int64     `json:"price"`
	Stock     int32     `json:"stock"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newUpdateProductResponse(product db.Product) updateProductResponse {
	return updateProductResponse{
		Name:      product.Name,
		Slug:      product.Slug,
		Sku:       product.Sku,
		Price:     product.Price,
		Stock:     product.Stock,
		UpdatedAt: product.UpdatedAt,
	}
}

type updateProductRequest struct {
	CategoryID int64  `json:"category_id" binding:"required,min=1"`
	ID         int64  `json:"id" binding:"required,min=1"`
	Name       string `json:"name" binding:"required"`
	Slug       string `json:"slug" binding:"required"`
	Sku        string `json:"sku" binding:"required"`
	Price      int64  `json:"price" binding:"required,min=1"`
	Stock      int32  `json:"stock" binding:"required,min=1"`
}

func (server *Server) updateProduct(ctx *gin.Context) {
	var req updateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.UpdateProductDTO{
		CategoryID: req.CategoryID,
		ID:         req.ID,
		Name:       req.Name,
		Slug:       req.Slug,
		Sku:        req.Sku,
		Price:      req.Price,
		Stock:      req.Stock,
	}
	product, err := server.productService.UpdateProduct(ctx, dto)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newUpdateProductResponse(product)
	ctx.JSON(http.StatusOK, rsp)
}

type updateProductStockRequest struct {
	ID    int64 `json:"id" binding:"required,min=1"`
	Stock int32 `json:"stock" binding:"required,min=1"`
}

func (server *Server) updateProductStock(ctx *gin.Context) {
	var req updateProductStockRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	dto := service.UpdateProductStockDTO{
		ID:    req.ID,
		Stock: req.Stock,
	}
	product, err := server.productService.UpdateProductStock(ctx, dto)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newUpdateProductResponse(product)
	ctx.JSON(http.StatusOK, rsp)
}

type deleteProductRequest struct {
	ID int64 `json:"id" binding:"required,min=1"`
}

func (server *Server) deleteProduct(ctx *gin.Context) {
	var req deleteProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.DeleteProductDTO{
		ID: req.ID,
	}
	product, err := server.productService.DeleteProduct(ctx, dto)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "deleted successful", "data": product})
}
