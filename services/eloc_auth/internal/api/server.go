package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_auth/db/sqlc"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_auth/internal/service"
)

type Server struct {
	store       db.Store
	userService *service.UserService
	router      *gin.Engine
}

func NewServer(store db.Store) (*Server, error) {
	userService := service.NewUserService(store)
	router := gin.Default()

	server := &Server{
		store:       store,
		userService: userService,
		router:      router,
	}
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {

	userRouters := server.router.Group("/users")
	{
		userRouters.POST("", server.CreateUser)
		userRouters.POST("/login", server.loginUser)
		userRouters.DELETE("/delete", server.deleteUser)
	}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
