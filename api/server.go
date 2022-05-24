package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/sheyi103/agtMiddleware/db/sqlc"
	"github.com/sheyi103/agtMiddleware/token"
	"github.com/sheyi103/agtMiddleware/util"
)

// Server serves HTTP requests for our AGT MIDDLEWARE
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker

	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setUpRouter()

	return server, nil
}

func (server *Server) setUpRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/roles", server.createRole)
	router.GET("/users/:id", server.getUser)
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	//authRoutes.GET("/users/:id", server.getUser)
	authRoutes.GET("/users", server.listUser)

	//authRoutes.POST("/roles", server.createRole)
	authRoutes.GET("/roles/:id", server.getRole)
	authRoutes.GET("/roles", server.listRole)

	authRoutes.POST("/services", server.createService)
	authRoutes.GET("/services/:id", server.getService)
	authRoutes.GET("/services", server.listService)

	authRoutes.POST("/shortcode", server.createShortCode)
	authRoutes.GET("/shortcode/:id", server.getShortCode)
	authRoutes.GET("/shortcode", server.listShortCodes)

	authRoutes.POST("/messages/sms/outbound", server.sendSMS)
	authRoutes.POST("/messages/ussd/outbound", server.sendUSSD)

	server.router = router
}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
