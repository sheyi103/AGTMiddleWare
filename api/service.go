package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	db "github.com/sheyi103/agtMiddleware/db/sqlc"
)

type ServicesServiceInterface struct {
	Name string `json:"service_interface" binding:"required,oneof=SOAP HTTP SMPP"`
}

type Service struct {
	Name string `json:"service" binding:"required,oneof=SMS USSD VOICE SUBCRIPTION"`
}

type createServiceRequest struct {
	ClientID                string                      `json:"client_id" binding:"required"`
	ClientSecret            string                      `json:"client_secret" binding:"required"`
	ShortcodeID             int32                       `json:"shortcode" binding:"required"`
	UserID                  int32                       `json:"user_id" binding:"required"`
	RoleID                  int32                       `json:"role_id" binding:"required"`
	ServiceName             string                      `json:"service_name"`
	ServiceID               string                      `json:"service_id"`
	ServiceInterface        db.ServicesServiceInterface `json:"service_interface" binding:"required,oneof=SOAP HTTP SMPP"`
	Service                 db.ServicesService          `json:"service" binding:"required,oneof=SMS USSD VOICE SUBCRIPTION"`
	ServiceType             db.ServicesServiceType      `json:"service_type" binding:"required,oneof=DAILY WEEKLY MONTHLY ON-DEMAND"`
	ProductID               string                      `json:"product_id"`
	NodeID                  string                      `json:"node_id"`
	SubscriptionID          string                      `json:"subscription_id"`
	SubscriptionDescription string                      `json:"subscription_description"`
	BaseUrl                 string                      `json:"base_url"`
	DatasyncEndpoint        string                      `json:"datasync_endpoint"`
	NotificationEndpoint    string                      `json:"notification_endpoint"`
	NetworkType             db.ServicesNetworkType      `json:"network_type" binding:"required,oneof=MTN AIRTEL GLO 9MOBILE"`
}

func (server *Server) createService(ctx *gin.Context) {
	var req createServiceRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//check if user exist if true get the client id and secret and add use to create a new service
	//Else create a new service with random generated string

	arg := db.CreateServiceParams{
		ClientID:                req.ClientID,
		ClientSecret:            req.ClientSecret,
		ShortcodeID:             req.ShortcodeID,
		UserID:                  req.UserID,
		RoleID:                  req.RoleID,
		ServiceName:             req.ServiceName,
		ServiceID:               req.ServiceID,
		ServiceInterface:        req.ServiceInterface,
		Service:                 req.Service,
		ServiceType:             req.ServiceType,
		ProductID:               req.ProductID,
		NodeID:                  req.NodeID,
		SubscriptionID:          req.SubscriptionID,
		SubscriptionDescription: req.SubscriptionDescription,
		BaseUrl:                 req.BaseUrl,
		DatasyncEndpoint:        req.DatasyncEndpoint,
		NotificationEndpoint:    req.NotificationEndpoint,
		NetworkType:             req.NetworkType,
	}

	_, err := server.store.CreateService(ctx, arg)

	if err != nil {

		if mysqlErr, ok := err.(*mysql.MySQLError); ok {

			switch mysqlErr.Number {
			case 1452, 1062:
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}

		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "Service created successfully")
}

type getServiceRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getService(ctx *gin.Context) {
	var req getUserRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	service, err := server.store.GetService(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return

		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, service)
}

type listServiceRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listService(ctx *gin.Context) {
	var req listServiceRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListServiceParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	services, err := server.store.ListService(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, services)
}
