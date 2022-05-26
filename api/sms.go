package api

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	db "github.com/sheyi103/agtMiddleware/db/sqlc"
	"github.com/sheyi103/agtMiddleware/madapi"
	"github.com/sheyi103/agtMiddleware/token"
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
	TargetSystem  string `json:"target_system" binding:"required"`
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
	//use the token to query for the users id
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	userId, err := server.store.GetUserByUsername(ctx, authPayload.Username)

	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	//once you have the users_id use it to query the service account
	args := db.GetServiceByUserIdParams{
		UserID:  userId.ID,
		Service: "SMS",
	}
	service, err := server.store.GetServiceByUserId(ctx, args)

	if err != nil {

		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	//update the notify url with the request notify url
	updateargs := db.UpdateNotifyEndpointByIdParams{
		NotificationEndpoint: req.NotifyUrl,
		ID:                   service.ID,
	}

	_, err = server.store.UpdateNotifyEndpointById(ctx, updateargs)

	if err != nil {

		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	//get the agt notify url from env and send it to madapi subscription

	//call sms subscription service
	smsSubscription, err := madapi.SMSSubscription(accessToken, req.SenderAddress, server.config.AGT_NOTIFY_URL, req.TargetSystem)
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

	//query the service table using the receiverAddress(shortcode )	where service is SMS to get notify url
	shortcodeId, err := server.store.GetShortcodeByShortCode(ctx, req.ReceiverAddress)
	if err != nil {

		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	log.Println(shortcodeId)

	notifyEndpoint, err := server.store.GetServiceByShortcodeId(ctx, shortcodeId)
	if err != nil {

		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	log.Println(notifyEndpoint.NotificationEndpoint)
	//forward traffic to the endpoint

	//call sms NotifyURl service
	_, err = madapi.SMSNotifyUrl(req.SenderAddress, req.ReceiverAddress, req.Message, req.Created, notifyEndpoint.NotificationEndpoint)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)

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
