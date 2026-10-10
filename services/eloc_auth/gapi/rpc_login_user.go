package gapi

import (
	"context"
	"database/sql"
	"errors"
	"time"

	pb "github.com/ryannguyen1105/eloc-backend/api/pb/eloc_auth"
	"github.com/ryannguyen1105/eloc-backend/common/util"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_auth/db/sqlc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := server.store.GetUserByEmail(ctx, db.GetUserByEmailParams{
		Email: req.GetEmail(),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to find user")
	}
	err = util.CheckPasswordHash(req.Password, user.PasswordHash)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "incorect password")
	}

	accessToken, err := server.tokenMaker.CreateToken(
		user.ID,
		user.Email,
		user.RoleID,
		user.IsVerified,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token")
	}

	refreshToken, err := server.tokenMaker.CreateToken(
		user.ID,
		user.Email,
		user.RoleID,
		user.IsVerified,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refresh token")
	}
	_, err = server.store.CreateUserToken(ctx, db.CreateUserTokenParams{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiresAt:    time.Now().Add(server.config.RefreshTokenDuration),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create session")
	}

	rsp := &pb.LoginUserResponse{
		User:         convertUser(user),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return rsp, nil
}
