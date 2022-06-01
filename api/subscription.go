package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sheyi103/agtMiddleware/madapi"
)

type dataSyncRequest struct {
	ServiceType   string `json:"serviceType" binding:"required"`
	ChargingMode  string `json:"chargingMode" binding:"required"`
	AppliedPlan   string `json:"appliedPlan" binding:"required"`
	ContentId     string `json:"contentId" binding:"required"`
	ResultCode    string `json:"resultCode" binding:"required"`
	RenFlag       string `json:"renFlag" binding:"required"`
	Result        string `json:"result" binding:"required"`
	ValidityType  string `json:"validityType" binding:"required"`
	SequenceNo    string `json:"sequenceNo" binding:"required"`
	CallingParty  string `json:"callingParty" binding:"required"`
	BearerId      string `json:"bearerId" binding:"required"`
	OperationId   string `json:"operationId" binding:"required"`
	RequestedPlan string `json:"requestedPlan" binding:"required"`
	ChargeAmount  string `json:"chargeAmount" binding:"required"`
	ServiceNode   string `json:"serviceNode" binding:"required"`
	ServiceId     string `json:"serviceId" binding:"required"`
	Category      string `json:"category" binding:"required"`
	ValidityDays  string `json:"validityDays" binding:"required"`
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

	var req dataSyncRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//get the shortcode

	//call sms subscription service
	// smsSubscription, err := madapi.SMSSubscription(accessToken, req.SenderAddress, req.NotifyUrl, req.TargetSystem)
	// if err != nil {

	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }

	ctx.JSON(http.StatusOK, req)

}

