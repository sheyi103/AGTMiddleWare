package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/sheyi103/agtMiddleware/util"
	"github.com/stretchr/testify/require"
)

func createRandomService(t *testing.T) int64 {
	role_id := int32(createRandomRole(t))
	user_id := int32(createRandomUser(t))
	shortcode_id := int32(createRandomShortCode(t))

	arg := CreateServiceParams{
		ClientID:                util.RandomString(15),
		ClientSecret:            util.RandomString(20),
		ShortcodeID:             shortcode_id,
		UserID:                  user_id,
		RoleID:                  role_id,
		ServiceName:             sql.NullString{String: util.RandomString(10), Valid: true},
		ServiceID:               sql.NullString{String: util.RandomString(10), Valid: true},
		ServiceInterface:        ServicesServiceInterface(util.RandomProtocol()),
		Service:                 ServicesService(util.RandomService()),
		ServiceType:             ServicesServiceType(util.RandomServiceType()),
		ProductID:               sql.NullString{String: util.RandomString(10), Valid: true},
		NodeID:                  sql.NullString{String: util.RandomString(10), Valid: true},
		SubscriptionID:          sql.NullString{String: util.RandomString(10), Valid: true},
		SubscriptionDescription: sql.NullString{String: util.RandomString(10), Valid: true},
		BaseUrl:                 sql.NullString{String: util.BaseUrl(), Valid: true},
		DatasyncEndpoint:        sql.NullString{String: util.DataSyncUrl(), Valid: true},
		NotificationEndpoint:    sql.NullString{String: util.NotificationUrl(), Valid: true},
		NetworkType:             ServicesNetworkType(util.RandomNetworkType()),
	}

	service_id, err := testQueries.CreateService(context.Background(), arg)
	require.NoError(t, err)
	rows, err := service_id.RowsAffected()
	require.NoError(t, err)
	require.Equal(t, int64(1), rows)
	//checking for not zero
	insertID, err := service_id.LastInsertId()
	require.NoError(t, err)
	require.NotZero(t, insertID)
	return insertID
}
func TestCreateService(t *testing.T) {
	createRandomService(t)
}

func TestGetService(t *testing.T) {
	//create service1
	service1 := createRandomService(t)
	//get service1 and assing it to service2 check for errors
	service2, err := testQueries.GetService(context.Background(), int32(service1))
	//check for errors
	require.NoError(t, err)
	//check if service2 is empty
	require.NotEmpty(t, service2)
	//eck that service1 is equal to service2
	require.Equal(t, int32(service1), service2.ID)
}
