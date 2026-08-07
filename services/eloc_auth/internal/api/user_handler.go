package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/ryannguyen1105/eloc-backend/common/middleware"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_auth/db/sqlc"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_auth/internal/service"
	"golang.org/x/crypto/bcrypt"
)

type userResponse struct {
	Email      string    `json:"email"`
	FullName   string    `json:"full_name"`
	IsActive   bool      `json:"is_active"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	Role       string    `json:"role"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Email:      user.Email,
		FullName:   user.Fullname,
		IsActive:   user.IsActive,
		IsVerified: user.IsVerified,
		CreatedAt:  user.CreatedAt,
		Role:       user.RoleID,
	}
}

type createUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Role     string `json:"role" binding:"required,uppercase,oneof=ADMIN STAFF CUSTOMER"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.CreateUserDTO{
		Email:    req.Email,
		Password: req.Password,
		FullName: req.FullName,
		Role:     req.Role,
	}

	user, err := server.authService.CreateUser(ctx, dto)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusConflict, errorResponse(errors.New("email already exists")))
				return

			case "foreign_key_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(errors.New("selected role does not exist")))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type loginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	Token string       `json:"token"`
	User  userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.LoginUserDTO{
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := server.authService.LoginUser(ctx, dto)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	token, err := server.tokenMaker.CreateToken(user.ID, user.Email, user.RoleID, user.IsVerified, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	_, err = server.store.CreateUserToken(ctx, db.CreateUserTokenParams{
		UserID:       user.ID,
		RefreshToken: token,
		ExpiresAt:    time.Now().Add(server.config.AccessTokenDuration),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := loginUserResponse{
		Token: token,
		User:  newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}

type updatedUserResponse struct {
	Email      string    `json:"email"`
	FullName   string    `json:"full_name"`
	IsActive   bool      `json:"is_active"`
	IsVerified bool      `json:"is_verified"`
	UpdatedAt  time.Time `json:"updated_at"`
	Role       string    `json:"role"`
}

func newUpdatedUserResponse(user db.User) updatedUserResponse {
	return updatedUserResponse{
		Email:      user.Email,
		FullName:   user.Fullname,
		IsActive:   user.IsActive,
		IsVerified: user.IsVerified,
		UpdatedAt:  user.UpdatedAt,
		Role:       user.RoleID,
	}
}

type updateUserFullNameRequest struct {
	OldName  string `json:"old_name" binding:"required"`
	NewName  string `json:"new_name" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

func (server *Server) updateUserFullName(ctx *gin.Context) {
	payload, err := middleware.GetAuthPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	var req updateUserFullNameRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.UpdateUserFullnameDTO{
		Email:    payload.Email,
		OldName:  req.OldName,
		NewName:  req.NewName,
		Password: req.Password,
	}
	updatedUser, err := server.authService.UpdateUserFullname(ctx, dto)
	if err != nil {
		if errors.Is(err, service.ErrOldNameMismatch) || errors.Is(err, service.ErrSameName) {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newUpdatedUserResponse(updatedUser)
	ctx.JSON(http.StatusOK, rsp)

}

type updateUserPasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required,min=6"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

func (server *Server) updateUserPassword(ctx *gin.Context) {
	payload, err := middleware.GetAuthPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	var req updateUserPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	dto := service.UpdateUserPasswordDTO{
		Email:       payload.Email,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}
	updatedUser, err := server.authService.UpdateUserPassword(ctx, dto)
	if err != nil {
		if errors.Is(err, service.ErrSamePassword) {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newUpdatedUserResponse(updatedUser)
	ctx.JSON(http.StatusOK, rsp)
}
