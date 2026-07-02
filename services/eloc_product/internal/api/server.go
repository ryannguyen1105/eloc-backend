package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_product/db/sqlc"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_product/internal/service"
)

type Server struct {
	store          db.Store
	productService *service.ProductService
	router         *gin.Engine
}

func NewServer(store db.Store) (*Server, error) {
	productService := service.NewProductService(store)
	router := gin.Default()

	server := &Server{
		store:          store,
		productService: productService,
		router:         router,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	categoryRouters := server.router.Group("/category")
	{
		categoryRouters.POST("", server.CreateCategory)
		categoryRouters.GET("",server.GetCategory)
		categoryRouters.DELETE("/delete", server.DeleteCategory)
	}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
