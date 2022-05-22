package db

import (
	"context"
	"log"
	"testing"

	"github.com/sheyi103/agtMiddleware/util"
	"github.com/stretchr/testify/require"
)

func createRandomService(t *testing.T) int64 {
	
	user_id := int32(createRandomUser(t))
	shortcode_id := int32(createRandomShortCode(t))

	arg := CreateServiceParams{
		ShortcodeID:             shortcode_id,
		UserID:                  user_id,
		ServiceName:             util.RandomString(10),
		ServiceID:               util.RandomString(10),
		ServiceInterface:        ServicesServiceInterface(util.RandomProtocol()),
		Service:                 ServicesService(util.RandomService()),
		ServiceType:             ServicesServiceType(util.RandomServiceType()),
		ProductID:                util.RandomString(10),
		NodeID:                   util.RandomString(10),
		SubscriptionID:           util.RandomString(10),
		SubscriptionDescription:  util.RandomString(10),
		BaseUrl:                  util.BaseUrl(),
		DatasyncEndpoint:         util.DataSyncUrl(), 
		NotificationEndpoint:     util.NotificationUrl(),
		NetworkType:             ServicesNetworkType(util.RandomNetworkType()),
	}

	service_id, err := testQueries.CreateService(context.Background(), arg)
	log.Print(err)
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
