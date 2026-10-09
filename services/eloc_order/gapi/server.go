package gapi

import (
	"github.com/gin-gonic/gin"
	pb "github.com/ryannguyen1105/eloc-backend/api/pb/eloc_order"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_order/config"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_order/db/sqlc"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_order/internal/service"
)

type Server struct {
	pb.UnimplementedEloc_OrderServer
	config       config.Config
	store        db.Store
	orderService *service.OrderService
	router       *gin.Engine
}

func NewServer(config config.Config, store db.Store) (*Server, error) {
	orderService := service.NewOrderService(store)
	router := gin.Default()

	server := &Server{
		config:       config,
		store:        store,
		orderService: orderService,
		router:       router,
	}

	return server, nil
}
