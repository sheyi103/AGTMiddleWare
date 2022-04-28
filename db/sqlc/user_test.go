package db

import (
	"context"
	"testing"

	"github.com/sheyi103/agtMiddleware/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) int64 {
	role_id := createRandomRole(t)
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Name:          util.RandomString(11),
		Email:         util.RandomEmail(),
		PhoneNumber:   util.RandomPhoneNumber(),
		Password:      hashedPassword,
		ContactPerson: util.RandomString(11),
		RoleID:        int32(role_id),
	}

	user_id, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	rows, err := user_id.RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(1), rows)
	//checking for not zero
	insertID, err := user_id.LastInsertId()
	require.NoError(t, err)
	require.NotZero(t, insertID)
	return insertID
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	//Get users
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), int32(user1))
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, int32(user1), user2.ID)
}
