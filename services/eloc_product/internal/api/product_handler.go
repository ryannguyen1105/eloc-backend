package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_product/db/sqlc"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_product/internal/service"
	"github.com/sqlc-dev/pqtype"
)

type createProductRequest struct {
	Name  string `json:"name" binding:"required"`
	Slug  string `json:"slug" binding:"required"`
	Sku   string `json:"sku" binding:"required"`
	Price int64  `json:"price" binding:"required,min=1"`
	Stock int32  `json:"stock" binding:"required,min=1"`
}

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
	Name string `json:"name" binding:"required"`
}

func (server *Server) getProduct(ctx *gin.Context) {
	var req getProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.GetProductDTO{
		Name: req.Name,
	}
	product, err := server.productService.GetProduct(ctx, dto)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newProductResponse(product)
	ctx.JSON(http.StatusOK, rsp)
}

type updateProductRequest struct {
	Name  string `json:"name" binding:"required"`
	Slug  string `json:"slug" binding:"required"`
	Sku   string `json:"sku" binding:"required"`
	Price int64  `json:"price" binding:"required,min=1"`
	Stock int32  `json:"stock" binding:"required,min=1"`
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

func (server *Server) updateProduct(ctx *gin.Context) {
	var req updateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	product, err := server.store.GetProductByName(ctx, db.GetProductByNameParams{
		Name: req.Name,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	product, err = server.store.UpdateProduct(ctx, db.UpdateProductParams{
		ID:         product.ID,
		CategoryID: product.CategoryID,
		Name:       req.Name,
		Slug:       req.Slug,
		Sku:        req.Sku,
		Price:      req.Price,
		Stock:      req.Stock,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
	}
	rsp := newProductResponse(product)
	ctx.JSON(http.StatusOK, rsp)
}

type updateProductStockRequest struct {
	Name  string `json:"name" binding:"required"`
	Stock int32  `json:"stock" binding:"required,min=1"`
}

func (server *Server) updateProductStock(ctx *gin.Context) {
	var req updateProductStockRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	product, err := server.store.GetProductByName(ctx, db.GetProductByNameParams{
		Name: req.Name,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	product, err = server.store.UpdateProductStock(ctx, db.UpdateProductStockParams{
		Name:  req.Name,
		Stock: req.Stock,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
	}
	rsp := newUpdateProductResponse(product)
	ctx.JSON(http.StatusOK, rsp)
}

type deleteProductRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) deleteProduct(ctx *gin.Context) {
	var req deleteProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	product, err := server.store.GetProductByName(ctx, db.GetProductByNameParams{
		Name: req.Name,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	err = server.store.DeleteProduct(ctx, db.DeleteProductParams{
		Name: req.Name,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "deleted successful", "data": product})
}
