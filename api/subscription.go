package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/sheyi103/agtMiddleware/db/sqlc"
	"github.com/sheyi103/agtMiddleware/madapi"
)

type dataSyncRequest struct {
	ServiceType   string `json:"serviceType"`
	ChargingMode  string `json:"chargingMode"`
	AppliedPlan   string `json:"appliedPlan"`
	ContentId     string `json:"contentId"`
	ResultCode    string `json:"resultCode"`
	RenFlag       string `json:"renFlag"`
	Result        string `json:"result"`
	ValidityType  string `json:"validityType"`
	SequenceNo    string `json:"sequenceNo"`
	CallingParty  string `json:"callingParty"`
	BearerId      string `json:"bearerId"`
	OperationId   string `json:"operationId"`
	RequestedPlan string `json:"requestedPlan"`
	ChargeAmount  string `json:"chargeAmount"`
	ServiceNode   string `json:"serviceNode"`
	ServiceId     string `json:"serviceId"`
	Category      string `json:"category"`
	ValidityDays  string `json:"validityDays"`
}

type subscriptionRequest struct {
	SubscriptionProviderId  string `json:"subscriptionProviderId"`
	NodeId                  string `json:"nodeId"`
	SubscriptionId          string `json:"subscriptionId"`
	SubscriptionDescription string `json:"subscriptionDescription"`
	RegistrationChannel     string `json:"registrationChannel"`
	CustomerId              string `json:"customerId"`
	TransactionId           string `json:"transactionId"`
}

type unSubscriptionRequest struct {
	SubscriptionProviderId  string `json:"subscriptionProviderId"`
	SubscriptionId          string `json:"subscriptionId"`
	SubscriptionDescription string `json:"subscriptionDescription"`
	CustomerId              string `json:"customerId"`
	TransactionId           string `json:"transactionId"`
}

func (server *Server) customerSubscription(ctx *gin.Context) {
	var req subscriptionRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//call sms service
	sub, err := madapi.Subscription(server.config.MADAPI_CLIENT_ID, req.SubscriptionProviderId, req.NodeId, req.SubscriptionId, req.SubscriptionDescription, req.RegistrationChannel, req.CustomerId, req.TransactionId)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sub)

}

func (server *Server) customerUnSubscription(ctx *gin.Context) {
	var req unSubscriptionRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//call sms service
	sendSMS, err := madapi.UnSubscription(server.config.MADAPI_CLIENT_ID, req.SubscriptionProviderId, req.SubscriptionId, req.SubscriptionDescription, req.CustomerId, req.TransactionId)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sendSMS)

}

func (server *Server) dataSync(ctx *gin.Context) {
	// log.Println(ctx.Data())

	var req dataSyncRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//check if service and product_id
	args := db.GetServiceByServiceIdAndProductIdParams{
		SubscriptionID:          req.ServiceId,
		SubscriptionDescription: req.RequestedPlan,
	}

	dataSyncUrl, err := server.store.GetServiceByServiceIdAndProductId(ctx, args)
	if err != nil {

		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	log.Println(dataSyncUrl)
	//forward request to datasync url

	respAgs := madapi.DataSyncRequestParams{
		ServiceType:   req.ServiceType,
		ChargingMode:  req.ChargingMode,
		AppliedPlan:   req.AppliedPlan,
		ContentId:     req.ContentId,
		ResultCode:    req.ResultCode,
		RenFlag:       req.RenFlag,
		Result:        req.Result,
		ValidityType:  req.ValidityType,
		SequenceNo:    req.SequenceNo,
		CallingParty:  req.CallingParty,
		BearerId:      req.BearerId,
		OperationId:   req.OperationId,
		RequestedPlan: req.RequestedPlan,
		ChargeAmount:  req.ChargeAmount,
		ServiceNode:   req.ServiceNode,
		ServiceId:     req.ServiceId,
		Category:      req.Category,
		ValidityDays:  req.ValidityDays,
	}

	_, err = madapi.DataSyncNotification(dataSyncUrl, respAgs)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)

}
