package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_product/db/sqlc"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_product/internal/service"
)

type categoryResponse struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func newCategoryResponse(category db.Category) categoryResponse {
	return categoryResponse{
		Name: category.Name,
		Slug: category.Slug,
	}
}

type createCategoryRequest struct {
	Name string `json:"name" binding:"required"`
	Slug string `json:"slug" binding:"required"`
}

func (server *Server) createCategory(ctx *gin.Context) {
	var req createCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.CreateCategoryDTO{
		Name: req.Name,
		Slug: req.Slug,
	}
	category, err := server.productService.CreateCategory(ctx, dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newCategoryResponse(category)
	ctx.JSON(http.StatusOK, rsp)
}

type getCategoryRequest struct {
	ID int64 `json:"id" binding:"required,min=1"`
}

func (server *Server) getCategory(ctx *gin.Context) {
	var req getCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.GetCategoryDTO{
		ID: req.ID,
	}
	category, err := server.productService.GetCategory(ctx, dto)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newCategoryResponse(category)
	ctx.JSON(http.StatusOK, rsp)
}

type deleteCategoryRequest struct {
	ID int64 `json:"id" binding:"required"`
}

func (server *Server) deleteCategory(ctx *gin.Context) {
	var req deleteCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.DeleteCategoryDTO{
		ID: req.ID,
	}
	category, err := server.productService.DeleteCategory(ctx, dto)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message": "deleted successful", "data": category})
}