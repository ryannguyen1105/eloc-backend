package api

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
		router:      router,
	}
	server.setupRouter()

	return server
}

func (server *Server) setupRouter() {
	roleRouters := server.router.Group("/roles")
	{
		roleRouters.POST("", server.CreateRole)
		roleRouters.GET("/:id", server.GetRole)
		roleRouters.GET("", server.ListRoles)
		roleRouters.PUT("/:id", server.UpdateRole)
		roleRouters.DELETE("/:id", server.DeleteRole)
	}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
