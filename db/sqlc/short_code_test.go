package db

import (
	"context"
	"strconv"
	"testing"

	"github.com/sheyi103/agtMiddleware/util"
	"github.com/stretchr/testify/require"
)

func createRandomShortCode(t *testing.T) int64 {
	code := int(util.RandomInt(3, 1000))
	shortcode := strconv.Itoa(code)

	shortCode, err := testQueries.CreateShortCode(context.Background(), shortcode)
	require.NoError(t, err)
	shortCodes, err := shortCode.RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(1), shortCodes)
	//checking for not zero
	insertID, err := shortCode.LastInsertId()
	require.NoError(t, err)
	require.NotZero(t, insertID)
	return insertID
}
func TestCreateShortCode(t *testing.T) {
	createRandomShortCode(t)
}

func TestGetShortCode(t *testing.T) {
	//Get role
	shortcode1 := createRandomShortCode(t)
	// log.Fatal(role1)
	shortcode2, err := testQueries.GetShortCode(context.Background(), int32(shortcode1))
	// log.Fatal(err)
	require.NoError(t, err)
	require.NotEmpty(t, shortcode2)
	require.Equal(t, int32(shortcode1), shortcode2.ID)
}
