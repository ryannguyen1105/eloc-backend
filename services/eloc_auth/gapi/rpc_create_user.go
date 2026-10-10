package gapi

import (
	"context"

	"github.com/lib/pq"
	pb "github.com/ryannguyen1105/eloc-backend/api/pb/eloc_auth"
	"github.com/ryannguyen1105/eloc-backend/common/util"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_auth/db/sqlc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}
	arg := db.CreateUserParams{
		Email:        req.GetEmail(),
		PasswordHash: hashedPassword,
		Fullname:     req.GetFullName(),
		RoleID:       req.GetRole(),
	}
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				status.Error(codes.AlreadyExists, err.Error())

			case "foreign_key_violation":
				status.Error(codes.PermissionDenied, err.Error())

			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user %s", err)
	}
	rsp := &pb.CreateUserResponse{
		User: convertUser(user),
	}
	return rsp, nil
}

