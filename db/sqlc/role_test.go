package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/sheyi103/agtMiddleware/util"
	"github.com/stretchr/testify/require"
)

func createRandomRole(t *testing.T) int64 {
	name := util.RandomRole()

	role, err := testQueries.CreateRole(context.Background(), name)
	require.NoError(t, err)
	rows, err := role.RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(1), rows)
	//checking for not zero
	insertID, err := role.LastInsertId()
	require.NoError(t, err)
	require.NotZero(t, insertID)
	return insertID
}
func TestCreateRole(t *testing.T) {
	createRandomRole(t)
}

func TestGetRole(t *testing.T) {
	//Get role
	role1 := createRandomRole(t)
	// log.Fatal(role1)
	role2, err := testQueries.GetRole(context.Background(), int32(role1))
	// log.Fatal(err)
	require.NoError(t, err)
	require.NotEmpty(t, role2)
	require.Equal(t, int32(role1), role2.ID)
}

func TestUpdateRole(t *testing.T) {
	role1 := createRandomRole(t)

	arg := UpdateRoleNamesParams{
		ID:   int32(role1),
		Name: util.RandomString(8),
	}

	role2, err := testQueries.UpdateRoleNames(context.Background(), arg)
	require.NoError(t, err)

	rows, err := role2.RowsAffected()
	require.NoError(t, err)
	require.NotZero(t, rows)
	require.Equal(t, int64(1), rows)

}

func TestListRoles(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomRole(t)
	}
	arg := ListRolesParams{
		Limit:  5,
		Offset: 5,
	}

	roles, err := testQueries.ListRoles(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, roles, 5)

	for _, role := range roles {
		require.NotEmpty(t, role)
	}
}

func TestDeleteAccount(t *testing.T) {
	role1 := createRandomRole(t)
	err := testQueries.DeleteRole(context.Background(), int32(role1))
	require.NoError(t, err)

	role2, err := testQueries.GetRole(context.Background(), int32(role1))
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, role2)
}
