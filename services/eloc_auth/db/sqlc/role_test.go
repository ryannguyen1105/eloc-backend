package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ryannguyen1105/eloc-backend/services/eloc_auth/util"
	"github.com/stretchr/testify/require"
)

func createRandomRole(t *testing.T) Role {
	arg := CreateRoleParams{
		Name: util.RandomName(),
	}
	role, err := testQueries.CreateRole(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, role)

	require.Equal(t, arg.Name, role.Name)

	require.NotZero(t, role.ID)
	return role
}

func TestCreateRole(t *testing.T) {
	createRandomRole(t)
}

func TestGetRole(t *testing.T) {
	role1 := createRandomRole(t)
	arg := GetRoleParams{
		ID: role1.ID,
	}
	role2, err := testQueries.GetRole(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, role2)

	require.Equal(t, role1.ID, role2.ID)
	require.Equal(t, role1.Name, role2.Name)
}

func TestUpdateRole(t *testing.T) {
	role1 := createRandomRole(t)

	updateArg := UpdateRoleParams{
		ID:   role1.ID,
		Name: util.RandomName(),
	}
	err := testQueries.UpdateRole(context.Background(), updateArg)

	require.NoError(t, err)

	arg := GetRoleParams {
		ID: role1.ID,
	}

	updateRole, err := testQueries.GetRole(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updateRole)
	require.Equal(t, updateArg.ID, updateRole.ID)
	require.Equal(t, updateArg.Name, updateRole.Name)
}

func TestDeleteRole(t *testing.T) {
	role1 := createRandomRole(t)
	
	deleteArg := DeleteRoleParams {
		ID: role1.ID,
	}

	err := testQueries.DeleteRole(context.Background(), deleteArg)
	require.NoError(t, err)

	arg := GetRoleParams {
		ID: role1.ID,
	}

	role2, err := testQueries.GetRole(context.Background(),arg )
	require.Error(t, err)
	require.ErrorIs(t, err, sql.ErrNoRows)

	require.Empty(t, role2)

}

func TestListRoles(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomRole(t)
	}
	arg := ListRolesParams {
		Limit: 5,
		Offset: 5,
	}
	roles, err := testQueries.ListRoles(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, roles, 5)

	for _, role := range roles {
		require.NotEmpty(t, role)
	}
}