package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/sheyi103/agtMiddleware/db/sqlc"
)

type createUserRequest struct {
	Name          string `json:"name" binding:"required"`
	Email         string `json:"email" binding:"required"`
	PhoneNumber   string `json:"phoneNumber" binding:"required"`
	ContactPerson string `json:"contactPerson" binding:"required"`
	RoleID        int32  `json:"roleID" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Name:          req.Name,
		Email:         req.Email,
		PhoneNumber:   req.PhoneNumber,
		ContactPerson: req.ContactPerson,
		RoleID:        req.RoleID,
	}

	_, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "user created successfully")
}
