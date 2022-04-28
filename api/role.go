package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createRoleRequest struct {
	Name string `json:"name" binding:"required,oneof=ADMINISTRATOR SP AGT"`
}

func (server *Server) createRole(ctx *gin.Context) {
	var req createRoleRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := req.Name

	_, err := server.store.CreateRole(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "role created successfully")
}
