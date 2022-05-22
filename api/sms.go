package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sheyi103/agtMiddleware/madapi"
)

type sendSMSRequest struct {
	ClientCorrelator string   `json:"clientCorrelator" binding:"required`
	Message          string   `json:"message" binding:"required`
	ReceiverAddress  []string `json:"receiverAddress" binding:"required`
	SenderAddress    string   `json:"senderAddress" binding:"required`
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
