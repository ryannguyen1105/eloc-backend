package gapi

import (
	pb "github.com/ryannguyen1105/eloc-backend/api/pb/eloc_auth"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_auth/db/sqlc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Email: user.Email,
		FullName: user.Fullname,
		IsActive: user.IsActive,
		IsVerified: user.IsVerified,
		Role: user.RoleID,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}