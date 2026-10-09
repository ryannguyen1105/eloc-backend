package gapi

import (
	"github.com/gin-gonic/gin"
	pb "github.com/ryannguyen1105/eloc-backend/api/pb/eloc_product"
	"github.com/ryannguyen1105/eloc-backend/common/middleware"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_product/config"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_product/db/sqlc"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_product/internal/service"
)

type Server struct {
	pb.UnimplementedEloc_ProductServer
	config            config.Config
	store             db.Store
	productService    *service.ProductService
	cloudinaryService *service.CloudinaryService
	router            *gin.Engine
}

func NewServer(config config.Config, store db.Store) (*Server, error) {
	router := gin.Default()
	productService := service.NewProductService(store)

	cloudinaryService, err := service.NewCloudinaryService()
	if err != nil {
		return nil, err
	}

	router.Use(middleware.CORSMiddleware())

	server := &Server{
		config:            config,
		store:             store,
		productService:    productService,
		router:            router,
		cloudinaryService: cloudinaryService,
	}

	return server, nil
}
