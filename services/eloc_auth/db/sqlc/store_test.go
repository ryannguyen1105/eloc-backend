package db

import (
	"context"
	"testing"
	"time"

	"github.com/ryannguyen1105/eloc-backend/util"
	"github.com/stretchr/testify/require"
)

func TestRegisterTx(t *testing.T) {
	store := NewStore(testDB)

	role := createRandomRole(t)
	user1 := createRandomUser(t, role)

	n := 5
	refreshToken := util.RandomRefreshToken()

	errs := make(chan error)
	results := make(chan RegisterTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.RegisterTx(context.Background(), RegisterTxParams{
				Email:        user1.Email,
				PasswordHash: user1.PasswordHash,
				FullName:     user1.Fullname,
				RoleID:       user1.RoleID,
				RefreshToken: refreshToken,
				ExpiresAt:    time.Now().UTC().Add(30 * time.Minute),
			})
			errs <- err
			results <- result
		}()
	}
	successCount := 0
	for i := 0; i < n; i++ {
		err := <-errs
		result := <-results

		if err == nil {
			successCount++
			require.NotEmpty(t, result.User)
			require.NotEmpty(t, result.Token)
			require.Equal(t, user1.Email, result.User.Email)
			require.Equal(t, refreshToken, result.Token.RefreshToken)
		} else {
			require.Error(t, err)
		}
	}
	require.Equal(t, 0, successCount)
}
