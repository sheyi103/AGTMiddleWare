package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/sheyi103/agtMiddleware/db/sqlc"
	"github.com/sheyi103/agtMiddleware/madapi"
	"github.com/sheyi103/agtMiddleware/token"
)

type sendUSSDRequest struct {
	SessionId   string `json:"sessionId" binding:"required`
	MessageType string `json:"messageType" binding:"required`
	Msisdn      string `json:"msisdn" binding:"required`
	ServiceCode string `json:"serviceCode" binding:"required`
	UssdString  string `json:"ussdString" binding:"required`
}

// type ussdNotifyRequest struct {
// 	SenderAddress   string `json:"senderAddress" binding:"required"`
// 	ReceiverAddress string `json:"receiverAddress" binding:"required"`
// 	Message         string `json:"message" binding:"required"`
// 	Created         int64  `json:"created" binding:"required"`
// }

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
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	userId, err := server.store.GetUserByUsername(ctx, authPayload.Username)

	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	//once you have the users_id use it to query the service account
	args := db.GetServiceByUserIdParams{
		UserID:  userId.ID,
		Service: "USSD",
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

	//call sms subscription service
	ussdSubscription, err := madapi.USSDSubscription(accessToken, req.SenderAddress, server.config.AGT_USSD_NOTIFY_URL, req.TargetSystem)
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

type USSDFlowResp struct {
	StatusCode string `json:"statusCode"`

	Data struct {
		InboundResponse   string `json:"inboundResponse"`
		UserInputRequired bool   `json:"userInputRequired"`
		MessageType       int32  `json:"messageType"`
		ServiceCode       string `json:"serviceCode"`
		Msisdn            string `json:"msisdn"`
		SessionId         string `json:"sessionId"`
	} `json:"data"`

	StatusMessage string `json:"statusMessage"`

	Link struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_link"`
}

func (server *Server) ussdNotify(ctx *gin.Context) {
	var req sendUSSDRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	shortcodeId, err := server.store.GetShortcodeByShortCode(ctx, req.ServiceCode)
	if err != nil {

		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	args := db.GetServiceByShortcodeIdParams{
		ShortcodeID: shortcodeId,
		Service:     "USSD",
	}

	notifyEndpoint, err := server.store.GetServiceByShortcodeId(ctx, args)
	if err != nil {

		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	log.Println(notifyEndpoint.NotificationEndpoint)
	//forward traffic to the endpoint

	//call sms NotifyURl service
	ussdResponse, err := madapi.USSDNotifyUrl(req.SessionId, req.MessageType, req.Msisdn, req.ServiceCode, req.UssdString, notifyEndpoint.NotificationEndpoint)
	log.Println("Inside controller")
	log.Println(ussdResponse)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ussdResponse)

}
