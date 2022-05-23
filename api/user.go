package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	db "github.com/sheyi103/agtMiddleware/db/sqlc"
	"github.com/sheyi103/agtMiddleware/util"
)

type createUserRequest struct {
	Name          string `json:"name" binding:"required"`
	Password      string `json:"password" binding:"required,min=6"`
	Email         string `json:"email" binding:"required,email"`
	PhoneNumber   string `json:"phone_number" binding:"required"`
	ContactPerson string `json:"contact_person" binding:"required"`
	RoleID        int32  `json:"role_id" binding:"required"`
}

type userResponse struct {
	Name          string `json:"name"`
	Password      string `json:"password"`
	ClientID      string `json:"client_id"`
	ClientSecret  string `json:"client_secret"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	ContactPerson string `json:"contact_person"`
	RoleID        int32  `json:"role_id"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Name:          user.Name,
		Password:      user.Password,
		ClientID:      user.ClientID,
		ClientSecret:  user.ClientSecret,
		Email:         user.Email,
		PhoneNumber:   user.PhoneNumber,
		ContactPerson: user.ContactPerson,
		RoleID:        user.RoleID,
		// CreatedAt:     user.CreatedAt,
		// UpdatedAt:     user.UpdatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Name:          req.Name,
		Password:      hashedPassword,
		ClientID:      util.RandomString(32),
		ClientSecret:  util.RandomString(16),
		Email:         req.Email,
		PhoneNumber:   req.PhoneNumber,
		ContactPerson: req.ContactPerson,
		RoleID:        req.RoleID,
	}

	_, err = server.store.CreateUser(ctx, arg)

	if err != nil {

		if mysqlErr, ok := err.(*mysql.MySQLError); ok {

			switch mysqlErr.Number {
			case 1452, 1062:
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}

		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "user created successfully")
}

type loginUserRequest struct {
	ClientID     string `json:"client_id" binding:"required,alphanum"`
	ClientSecret string `json:"client_secret" binding:"required,alphanum"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

type loginUser struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetUserByClientIdParams{

		ClientID:     req.ClientID,
		ClientSecret: req.ClientSecret,
	}

	user, err := server.store.GetUserByClientId(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(
		user.Email,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)

}

type getUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return

		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type lisUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listUser(ctx *gin.Context) {
	var req lisUserRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := server.store.ListUsers(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}
