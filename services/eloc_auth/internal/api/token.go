package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_auth/db/sqlc"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (server *Server) renewAccessToken(ctx *gin.Context) {
	var req renewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	session, err := server.store.GetUserTokenByRefreshToken(ctx, db.GetUserTokenByRefreshTokenParams{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if session.IsBlocked {
		err := fmt.Errorf("blocked session")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	if session.UserID != refreshPayload.UserID {
		err := fmt.Errorf("incorect session user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	if session.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("missmatched session token")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	if time.Now().After(session.ExpiresAt) {
		err := fmt.Errorf("expired session")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	user, err := server.store.GetUserByEmail(ctx, db.GetUserByEmailParams{
		Email: refreshPayload.Email,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	accessToken, err := server.tokenMaker.CreateToken(
		user.ID,
		user.Email,
		user.RoleID,
		user.IsVerified,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := renewAccessTokenResponse{
		AccessToken: accessToken,
		AccessTokenExpiresAt: time.Now().Add(server.config.AccessTokenDuration),
	}
	ctx.JSON(http.StatusOK, rsp)
}
