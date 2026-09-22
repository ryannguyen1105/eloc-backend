package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ryannguyen1105/eloc-backend/common/middleware"
	"github.com/ryannguyen1105/eloc-backend/common/token"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_auth/config"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_auth/db/sqlc"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_auth/internal/service"
)

type Server struct {
	config      config.Config
	store       db.Store
	tokenMaker  token.Maker
	authService *service.AuthService
	router      *gin.Engine
}

func NewServer(config config.Config, store db.Store) (*Server, error) {
	router := gin.Default()
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	authService := service.NewAuthService(store, tokenMaker, config)

	router.Use(middleware.CORSMiddleware())

	server := &Server{
		config:      config,
		store:       store,
		tokenMaker:  tokenMaker,
		authService: authService,
		router:      router,
	}
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {

	userRouters := server.router.Group("users")
	{
		userRouters.POST("", server.createUser)
		userRouters.POST("/login", server.loginUser)
		userRouters.POST("/token/renew_access", server.renewAccessToken)
	}
	authRouters := server.router.Group("/users").Use(middleware.AuthMiddleware(server.tokenMaker))
	{
		authRouters.PATCH("/updateFullName", server.updateUserFullName)
		authRouters.PATCH("/updatePassword", server.updateUserPassword)

	}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
