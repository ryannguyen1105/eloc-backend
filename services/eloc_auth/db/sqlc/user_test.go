package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ryannguyen1105/eloc-backend/services/eloc_auth/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T, role Role) User {
	hashPassword, err := util.HashPassword(util.RandomPassword())
	require.NoError(t, err)
	arg := CreateUserParams{
		Email:        util.RandomEmail(),
		PasswordHash: hashPassword,
		Fullname:     role.Name,
		RoleID:       role.ID,
		IsActive:     false,
		IsVerified:   false,
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)
	require.Equal(t, arg.Fullname, user.Fullname)
	require.Equal(t, arg.IsActive, user.IsActive)
	require.Equal(t, arg.IsVerified, user.IsVerified)

	return user
}

func TestCreateUser(t *testing.T) {
	role := createRandomRole(t)
	createRandomUser(t, role)
}

func TestGetUserByID(t *testing.T) {
	role := createRandomRole(t)
	user1 := createRandomUser(t, role)
	arg := GetUserByIdParams{
		ID: user1.ID,
	}
	user2, err := testQueries.GetUserById(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Fullname, user2.Fullname)
}

func TestGetUserByEmail(t *testing.T) {
	role := createRandomRole(t)
	user1 := createRandomUser(t, role)
	arg := GetUserByEmailParams{
		Email: user1.Email,
	}
	user2, err := testQueries.GetUserByEmail(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Fullname, user2.Fullname)

}

func TestUpdateUserDetail(t *testing.T) {
	role := createRandomRole(t)
	user1 := createRandomUser(t, role)
	arg := UpdateUserDetailParams{
		ID:       user1.ID,
		Fullname: user1.Fullname,
		RoleID:   role.ID,
	}
	user2, err := testQueries.UpdateUserDetail(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, arg.Fullname, user2.Fullname)
	require.Equal(t, arg.RoleID, user2.RoleID)
	require.Equal(t, user1.IsActive, user2.IsActive)
	require.Equal(t, user1.IsVerified, user2.IsVerified)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, time.Now(), user2.UpdatedAt, time.Second)

}


func TestUpdateUserStatus(t *testing.T) {
	role := createRandomRole(t)
	user1 := createRandomUser(t, role)
	arg := UpdateUserStatusParams {
		ID: user1.ID,
		IsActive: true,
		IsVerified: true,
	}
	err := testQueries.UpdateUserStatus(context.Background(), arg)
	require.NoError(t, err)
	
	getArg := GetUserByIdParams {
		ID: user1.ID,
	}
	user2, err := testQueries.GetUserById(context.Background(), getArg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.IsActive, user2.IsActive)
	require.Equal(t, arg.IsVerified, user2.IsVerified)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)

	require.WithinDuration(t, time.Now(), user1.UpdatedAt, time.Second)

}

func TestDeleteUser(t *testing.T) {
	role := createRandomRole(t)
	user1 := createRandomUser(t, role)

	deleteArg := DeleteUserParams {
		ID: user1.ID,
	}
	err := testQueries.DeleteUser(context.Background(), deleteArg)
	require.NoError(t, err)

	arg := GetUserByIdParams{
		ID: user1.ID,
	}
	user2, err := testQueries.GetUserById(context.Background(), arg)
	require.Error(t, err)
	require.ErrorIs(t, err, sql.ErrNoRows)

	require.Empty(t, user2)
}

func TestListUsers(t *testing.T) {
	role := createRandomRole(t)
	for i := 0; i < 5; i++ {
		createRandomUser(t, role)
	}
	arg := ListUsersParams {
		Limit: 5,
		Offset: 5,
	}
	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}