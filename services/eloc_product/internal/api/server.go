package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ryannguyen1105/eloc-backend/common/middleware"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_product/db/sqlc"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_product/internal/service"
)

type Server struct {
	store             db.Store
	productService    *service.ProductService
	cloudinaryService *service.CloudinaryService
	router            *gin.Engine
}

func NewServer(store db.Store) (*Server, error) {
	router := gin.Default()
	productService := service.NewProductService(store)

	cloudinaryService, err := service.NewCloudinaryService()
	if err != nil {
		return nil, err
	}

	router.Use(middleware.CORSMiddleware())

	server := &Server{
		store:             store,
		productService:    productService,
		router:            router,
		cloudinaryService: cloudinaryService,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	categoryRouters := server.router.Group("/category")
	{
		categoryRouters.POST("", server.createCategory)
		categoryRouters.GET("", server.getCategory)
		categoryRouters.DELETE("/delete", server.deleteCategory)
	}

	productsRouters := server.router.Group("/products")
	{
		productsRouters.GET("", server.ListProduct)
	}

	productRouters := server.router.Group("/product")
	{
		productRouters.POST("", server.createProduct)
		productRouters.GET("/:id", server.getProduct)
		productRouters.PUT("/update/:id", server.updateProduct)
		productRouters.PATCH("/updatestock/:id", server.updateProductStock)
		productRouters.DELETE("/delete", server.deleteProduct)

		productRouters.POST("/:product_id/images", server.addProductImage)
		productRouters.DELETE("/:id/:product_id", server.deleteProductImage)
	}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
