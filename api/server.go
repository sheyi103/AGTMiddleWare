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

	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	router.GET("/users/:id", server.getUser)
	router.GET("/users", server.listUser)

	router.POST("/roles", server.createRole)
	router.GET("/roles/:id", server.getRole)
	router.GET("/roles", server.listRole)

	router.POST("/services", server.createService)
	router.GET("/services/:id", server.getService)
	router.GET("/services", server.listService)

	router.POST("/shortcode", server.createShortCode)
	router.GET("/shortcode/:id", server.getShortCode)
	router.GET("/shortcode", server.listShortCodes)

	router.POST("/messages/sms/outbound", server.sendSMS)

	server.router = router
}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
