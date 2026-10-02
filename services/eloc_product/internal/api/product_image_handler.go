package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_product/db/sqlc"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_product/internal/service"
)

type productImageResponse struct {
	ID        int64  `json:"id"`
	ProductID int64  `json:"product_id"`
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func newProductImageResponse(productImage db.ProductImage) productImageResponse {
	return productImageResponse{
		ID:        productImage.ID,
		ProductID: productImage.ProductID,
		ImageUrl:  productImage.ImageUrl,
		IsPrimary: productImage.IsPrimary,
	}
}

type addProductImageRequest struct {
	ProductID int64 `uri:"product_id" binding:"required,min=1"`
	IsPrimary bool  `form:"is_primary" json:"is_primary"`
}

func (server *Server) addProductImage(ctx *gin.Context) {
	var req addProductImageRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer file.Close()

	imageUrl, err := server.cloudinaryService.UploadImage(ctx, file, "products")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	dto := service.AddProductImageDTO{
		ProductID: req.ProductID,
		ImageUrl:  imageUrl,
		IsPrimary: req.IsPrimary,
	}

	productImage, err := server.productService.AddProductImage(ctx, dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newProductImageResponse(productImage)
	ctx.JSON(http.StatusOK, rsp)

}

type deleteProductImageRequest struct {
	ID        int64 `uri:"id" binding:"required,min=1"`
	ProductID int64 `uri:"product_id" binding:"required,min=1"`
}

func (server *Server) deleteProductImage(ctx *gin.Context) {
	var req deleteProductImageRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.DeleteProductImageDTO{
		ID:        req.ID,
		ProductID: req.ProductID,
	}
	productImage, err := server.productService.DeleteProductImage(ctx, dto)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "delete successful", "data": productImage})
}
