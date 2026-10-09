package gapi

import (
	"fmt"

	pb "github.com/ryannguyen1105/eloc-backend/api/pb/eloc_auth"
	"github.com/gin-gonic/gin"
	"github.com/ryannguyen1105/eloc-backend/common/middleware"
	"github.com/ryannguyen1105/eloc-backend/common/token"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_auth/config"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_auth/db/sqlc"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_auth/internal/service"
)

type Server struct {
	pb.UnimplementedEloc_AuthServer
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

	return server, nil
}
