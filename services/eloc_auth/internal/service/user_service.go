package service

import (
	"context"

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

type UpdateUserDetailDTO struct {
	Email    string
	Password string
	FullName string
}

func (userService *AuthService) UpdateUserDetail(ctx context.Context, dto UpdateUserDetailDTO) (db.User, error) {
	user, err := userService.store.GetUserByEmail(ctx, db.GetUserByEmailParams{
		Email: dto.Email,
	})
	if err != nil {
		return db.User{}, nil
	}

	arg := db.UpdateUserDetailParams{
		ID:           user.ID,
		PasswordHash: user.PasswordHash,
		Fullname:     dto.FullName,
		RoleID:       user.RoleID,
	}
	return userService.store.UpdateUserDetail(ctx, arg)
}

type UpdateUserPasswordDTO struct {
	Email       string
	OldPassword string
	NewPassword string
}

func (userService *AuthService) UpdateUserPassword(ctx context.Context, dto UpdateUserPasswordDTO) (db.User, error) {
	user, err := userService.store.GetUserByEmail(ctx, db.GetUserByEmailParams{
		Email: dto.Email,
	})
	if err != nil {
		return db.User{}, nil
	}
	if err := util.CheckPasswordHash(dto.OldPassword, user.PasswordHash); err != nil {
		return db.User{}, nil
	}
	newHashedPassword, err := util.HashPassword(dto.NewPassword)
	if err != nil {
		return db.User{}, nil
	}
	arg := db.UpdateUserPasswordParams{
		ID:           user.ID,
		PasswordHash: newHashedPassword,
	}
	return userService.store.UpdateUserPassword(ctx, arg)
}

type DeleteUserDTO struct {
	Email    string
	Password string
}

func (userService *AuthService) DeleteUser(ctx context.Context, dto DeleteUserDTO) (db.User, error) {
	user, err := userService.store.GetUserByEmail(ctx, db.GetUserByEmailParams{
		Email: dto.Email,
	})
	if err != nil {
		return db.User{}, err
	}
	if err := util.CheckPasswordHash(dto.Password, user.PasswordHash); err != nil {
		return db.User{}, nil
	}
	err = userService.store.DeleteUser(ctx, db.DeleteUserParams{
		Email: user.Email,
	})
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}
