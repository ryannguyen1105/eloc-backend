package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ryannguyen1105/eloc-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomUserToken(t *testing.T, user User) UserToken {
	arg := CreateUserTokenParams{
		UserID:       user.ID,
		RefreshToken: util.RandomRefreshToken(),
		ExpiresAt:    time.Now().UTC().Add(30 * time.Minute),
	}
	userToken, err := testQueries.CreateUserToken(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, userToken)

	require.Equal(t, arg.UserID, userToken.UserID)
	require.Equal(t, arg.RefreshToken, userToken.RefreshToken)
	require.WithinDuration(t, arg.ExpiresAt, userToken.ExpiresAt, time.Second)

	return userToken
}

func TestCreateUserToken(t *testing.T) {
	role := createRandomRole(t)
	user := createRandomUser(t, role)

	createRandomUserToken(t, user)
}

func TestGetUserToken(t *testing.T) {
	role := createRandomRole(t)
	user := createRandomUser(t, role)

	userToken1 := createRandomUserToken(t, user)
	arg := GetUserTokenParams{
		RefreshToken: userToken1.RefreshToken,
	}
	userToken2, err := testQueries.GetUserToken(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userToken2)

	require.Equal(t, userToken1.ID, userToken2.ID)
	require.Equal(t, userToken1.RefreshToken, userToken2.RefreshToken)
	require.WithinDuration(t, userToken1.ExpiresAt, userToken2.ExpiresAt, time.Second)
}

func TestDeleteUserToken(t *testing.T) {
	role := createRandomRole(t)
	user := createRandomUser(t, role)

	userToken1 := createRandomUserToken(t, user)
	deleteArg := DeleteUserTokenParams{
		RefreshToken: userToken1.RefreshToken,
	}
	err := testQueries.DeleteUserToken(context.Background(), deleteArg)
	require.NoError(t, err)

	arg := GetUserTokenParams{
		RefreshToken: userToken1.RefreshToken,
	}
	userToken2, err := testQueries.GetUserToken(context.Background(), arg)
	require.Error(t, err)
	require.ErrorIs(t, err, sql.ErrNoRows)

	require.Empty(t, userToken2)
}

func TestDeleteAllUserToken(t *testing.T) {
	role := createRandomRole(t)
	user := createRandomUser(t, role)

	userToken1 := createRandomUserToken(t, user)
	deleteArg := DeleteAllUserTokensParams{
		UserID: userToken1.UserID,
	}
	err := testQueries.DeleteAllUserTokens(context.Background(), deleteArg)
	require.NoError(t, err)

	arg := GetUserTokenParams{
		RefreshToken: userToken1.RefreshToken,
	}
	userToken2, err := testQueries.GetUserToken(context.Background(), arg)
	require.Error(t, err)
	require.ErrorIs(t, err, sql.ErrNoRows)

	require.Empty(t, userToken2)
}

func TestDeleteExpiredTokens(t *testing.T) {
	role := createRandomRole(t)
	user := createRandomUser(t, role)

	expiredArg := CreateUserTokenParams{
		UserID:       user.ID,
		RefreshToken: util.RandomRefreshToken(),
		ExpiresAt:    time.Now().UTC().Add(-5 * time.Minute),
	}
	expiredToken, err := testQueries.CreateUserToken(context.Background(), expiredArg)
	require.NoError(t, err)
	require.NotEmpty(t, expiredToken)

	liveToken := createRandomUserToken(t, user)
	err = testQueries.DeleteExpiredTokens(context.Background())
	require.NoError(t, err)

	getArg1 := GetUserTokenParams{
		RefreshToken: expiredToken.RefreshToken,
	}
	res1, err := testQueries.GetUserToken(context.Background(), getArg1)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, res1)

	getArg2 := GetUserTokenParams{
		RefreshToken: liveToken.RefreshToken,
	}
	res2, err := testQueries.GetUserToken(context.Background(), getArg2)
	require.NoError(t, err)
	require.NotEmpty(t, res2)
	require.Equal(t, liveToken.RefreshToken, res2.RefreshToken)
}
