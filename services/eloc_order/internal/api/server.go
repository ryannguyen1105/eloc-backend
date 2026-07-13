package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_order/db/sqlc"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_order/internal/service"
)

type Server struct {
	store        db.Store
	orderService *service.OrderService
	router       *gin.Engine
}

func NewServer(store db.Store) (*Server, error) {
	orderService := service.NewOrderService(store)
	router := gin.Default()

	server := &Server{
		store:        store,
		orderService: orderService,
		router:       router,
	}
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	cartRouters := server.router.Group("/cart")
	{
		cartRouters.POST("", server.createCart)
		cartRouters.GET("", server.getCart)
		cartRouters.PATCH("/update", server.updateCartQuantity)
		cartRouters.DELETE("/remove", server.removeFromCart)
		cartRouters.DELETE("/clear", server.clearFromCart)
	}

	orderRouters := server.router.Group("/order")
	{
		orderRouters.POST("", server.createOrder)
		orderRouters.GET("", server.getOrder)
		orderRouters.GET("/list", server.listOrder)
		orderRouters.PATCH("/update", server.updateOrderStatus)
	}

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
