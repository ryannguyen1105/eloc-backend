package service

import (
	"context"
	"errors"

	"github.com/ryannguyen1105/eloc-backend/common/util"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_auth/db/sqlc"
)

type CreateUserDTO struct {
	Email    string
	Password string
	FullName string
	Role     string
}

func (userService *AuthService) CreateUser(ctx context.Context, dto CreateUserDTO) (db.User, error) {
	hashedPassword, err := util.HashPassword(dto.Password)
	if err != nil {
		return db.User{}, err
	}
	arg := db.CreateUserParams{
		Email:        dto.Email,
		PasswordHash: hashedPassword,
		Fullname:     dto.FullName,
		RoleID:       dto.Role,
		IsActive:     true,
	}
	return userService.store.CreateUser(ctx, arg)
}

type LoginUserDTO struct {
	Email    string
	Password string
}

func (userService *AuthService) LoginUser(ctx context.Context, dto LoginUserDTO) (db.User, error) {
	user, err := userService.store.GetUserByEmail(ctx, db.GetUserByEmailParams{
		Email: dto.Email,
	})
	if err != nil {
		return db.User{}, err
	}

	err = util.CheckPasswordHash(dto.Password, user.PasswordHash)
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}

type UpdateUserFullnameDTO struct {
	Email    string
	OldName  string
	NewName  string
	Password string
}

var (
	ErrOldNameMismatch = errors.New("old name does not match current name")
	ErrSameName        = errors.New("new name cannot be the same as current name")
)

func (userService *AuthService) UpdateUserFullname(ctx context.Context, dto UpdateUserFullnameDTO) (db.User, error) {
	user, err := userService.store.GetUserByEmail(ctx, db.GetUserByEmailParams{
		Email: dto.Email,
	})
	if err != nil {
		return db.User{}, err
	}
	if err := util.CheckPasswordHash(dto.Password, user.PasswordHash); err != nil {
		return db.User{}, err
	}
	if user.Fullname != dto.OldName {
		return db.User{}, ErrOldNameMismatch
	}

	if user.Fullname == dto.NewName {
		return db.User{}, ErrSameName
	}

	arg := db.UpdateUserFullnameParams{
		ID:       user.ID,
		Fullname: dto.NewName,
	}
	return userService.store.UpdateUserFullname(ctx, arg)
}

type UpdateUserPasswordDTO struct {
	Email       string
	OldPassword string
	NewPassword string
}

var ErrSamePassword = errors.New("new password cannot be the same as old password")

func (userService *AuthService) UpdateUserPassword(ctx context.Context, dto UpdateUserPasswordDTO) (db.User, error) {
	if dto.OldPassword == dto.NewPassword {
		return db.User{}, ErrSamePassword
	}

	user, err := userService.store.GetUserByEmail(ctx, db.GetUserByEmailParams{
		Email: dto.Email,
	})
	if err != nil {
		return db.User{}, err
	}
	if err := util.CheckPasswordHash(dto.OldPassword, user.PasswordHash); err != nil {
		return db.User{}, err
	}
	newHashedPassword, err := util.HashPassword(dto.NewPassword)
	if err != nil {
		return db.User{}, err
	}
	arg := db.UpdateUserPasswordParams{
		ID:           user.ID,
		PasswordHash: newHashedPassword,
	}
	return userService.store.UpdateUserPassword(ctx, arg)
}
