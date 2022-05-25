package api

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sheyi103/agtMiddleware/madapi"
)

type sendSMSRequest struct {
	ClientCorrelator string   `json:"clientCorrelator" binding:"required`
	Message          string   `json:"message" binding:"required`
	ReceiverAddress  []string `json:"receiverAddress" binding:"required`
	SenderAddress    string   `json:"senderAddress" binding:"required`
}

type smsDataSyncRequest struct {
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

type smsSubscriptionRequest struct {
	SenderAddress string `json:"sender_address" binding:"required"`
	NotifyUrl     string `json:"notify_url" binding:"required"`
	TargetSystem  string `json:"targetSystem" binding:"required"`
}

type smsNotifyRequest struct {
	SenderAddress   string `json:"senderAddress" binding:"required"`
	ReceiverAddress string `json:"receiverAddress" binding:"required"`
	Message         string `json:"message" binding:"required"`
	Created         int64  `json:"created" binding:"required"`
}

type smsDeleteSubscriptionRequest struct {
	SenderAddress  string `json:"sender_address" binding:"required"`
	SubscriptionId string `json:"subscriptionId" binding:"required"`
}

type authorizationResponse struct {
	AccessToken string `json:"accessToken"`
}

func (server *Server) sendSMS(ctx *gin.Context) {
	var req sendSMSRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	accessToken, err := madapi.Authorization(
		server.config.MADAPI_CLIENT_ID,
		server.config.MADAPI_CLIENT_SECRET,
	)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//call sms service
	sendSMS, err := madapi.SendOutBoundSMS(accessToken, req.ClientCorrelator, req.Message, req.ReceiverAddress, req.SenderAddress)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sendSMS)

}

func (server *Server) smsSubscription(ctx *gin.Context) {
	var req smsSubscriptionRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	accessToken, err := madapi.Authorization(
		server.config.MADAPI_CLIENT_ID,
		server.config.MADAPI_CLIENT_SECRET,
	)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//call sms subscription service
	smsSubscription, err := madapi.SMSSubscription(accessToken, req.SenderAddress, req.NotifyUrl, req.TargetSystem)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, smsSubscription)

}

func (server *Server) smsDeleteSubscription(ctx *gin.Context) {
	var req smsDeleteSubscriptionRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	accessToken, err := madapi.Authorization(
		server.config.MADAPI_CLIENT_ID,
		server.config.MADAPI_CLIENT_SECRET,
	)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//call sms subscription service
	smsSubscription, err := madapi.SMSDeleteSubscription(accessToken, req.SenderAddress, req.SubscriptionId)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, smsSubscription)

}

func (server *Server) SMSNotifyUrl(ctx *gin.Context) {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	var req smsNotifyRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//call sms subscription service
	// smsSubscription, err := madapi.SMSSubscription(accessToken, req.SenderAddress, req.NotifyUrl, req.TargetSystem)
	// if err != nil {

	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }

	ctx.JSON(http.StatusOK, req)

}

func (server *Server) dataSync(ctx *gin.Context) {

	var req smsDataSyncRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//call sms subscription service
	// smsSubscription, err := madapi.SMSSubscription(accessToken, req.SenderAddress, req.NotifyUrl, req.TargetSystem)
	// if err != nil {

	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }

	ctx.JSON(http.StatusOK, req)

}
