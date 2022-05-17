package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/sheyi103/agtMiddleware/db/sqlc"
)

// Server serves HTTP requests for our AGT MIDDLEWARE
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUser)
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

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
