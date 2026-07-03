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

type LoginUserDTO struct {
	Email    string
	Password string
}

type DeleteUserDTO struct {
	Email string
}

func (userService *AuthService) CreateUser(ctx context.Context, dto CreateUserDTO) (db.User, error) {
	role, err := userService.store.CreateRole(ctx, db.CreateRoleParams{
		ID:          dto.Role,
		Description: "",
	})
	if err != nil {
		return db.User{}, err
	}

	arg := db.CreateUserParams{
		Email:        dto.Email,
		PasswordHash: dto.Password,
		Fullname:     dto.FullName,
		RoleID:       role.ID,
		IsActive:     true,
	}
	return userService.store.CreateUser(ctx, arg)
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

func (userService *AuthService) DeleteUser(ctx context.Context, dto DeleteUserDTO) (db.User, error) {
	user, err := userService.store.GetUserByEmail(ctx, db.GetUserByEmailParams{
		Email: dto.Email,
	})
	if err != nil {
		return db.User{}, err
	}
	err = userService.store.DeleteUser(ctx,db.DeleteUserParams{
		Email: user.Email,
	})
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}
