package handler

import (
	"github.com/gin-gonic/gin"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_auth/db/sqlc"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_auth/internal/service"
)

type Server struct {
	roleService *service.RoleService
	router      *gin.Engine
}

func NewServer(store *db.Store) *Server {
	roleService := service.NewRoleService(store)
	router := gin.Default()

	server := &Server{
		roleService: roleService,
		router: router,
	}

	router.POST("/roles", server.CreateRole)
	router.GET("/roles/:id", server.GetRole)
	router.GET("/roles", server.ListRoles)
	router.PUT("/roles/:id", server.UpdateRole)
	router.DELETE("/roles/:id", server.DeleteRole)

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
