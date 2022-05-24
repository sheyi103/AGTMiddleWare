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
