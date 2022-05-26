package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sheyi103/agtMiddleware/madapi"
)

type sendUSSDRequest struct {
	SessionId   string `json:"session_id" binding:"required`
	MessageType string `json:"message_type" binding:"required`
	Msisdn      string `json:"msisdn" binding:"required`
	ServiceCode string `json:"service_code" binding:"required`
	UssdString  string `json:"ussd_string" binding:"required`
}

type ussdNotifyRequest struct {
	SenderAddress   string `json:"senderAddress" binding:"required"`
	ReceiverAddress string `json:"receiverAddress" binding:"required"`
	Message         string `json:"message" binding:"required"`
	Created         int64  `json:"created" binding:"required"`
}

type ussdSubscriptionRequest struct {
	SenderAddress string `json:"sender_address" binding:"required"`
	NotifyUrl     string `json:"notify_url" binding:"required"`
	TargetSystem  string `json:"target_System" binding:"required"`
}

type ussdDeleteSubscriptionRequest struct {
	SubscriptionId string `json:"subscriptionId" binding:"required"`
}

func (server *Server) sendUSSD(ctx *gin.Context) {
	var req sendUSSDRequest

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
	sendSMS, err := madapi.SendOutBoundUSSD(accessToken, req.SessionId, req.MessageType, req.Msisdn, req.ServiceCode, req.UssdString)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sendSMS)

}

func (server *Server) ussdSubscription(ctx *gin.Context) {
	var req ussdSubscriptionRequest

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
	//once you have the users_id use it to query the service account
	//update the notify url with the request notify url
	//get the agt notify url from env and send it to madapi subscription

	//call sms subscription service
	ussdSubscription, err := madapi.USSDSubscription(accessToken, req.SenderAddress, req.NotifyUrl, req.TargetSystem)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ussdSubscription)

}

func (server *Server) ussdDeleteSubscription(ctx *gin.Context) {
	var req ussdDeleteSubscriptionRequest

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
	ussdDeleteSubscription, err := madapi.USSDDeleteSubscription(accessToken, req.SubscriptionId)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ussdDeleteSubscription)

}

func (server *Server) USSDNotifyUrl(ctx *gin.Context) {
	var req sendUSSDRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	//query the database using the receiverAddress(shortcode )	where service is SMS
	//return the notify url that was updated earlier
	//forward traffic to the endpoint

	//call sms subscription service
	// smsSubscription, err := madapi.SMSSubscription(accessToken, req.SenderAddress, req.NotifyUrl, req.TargetSystem)
	// if err != nil {

	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }

	ctx.JSON(http.StatusOK, req)

}
