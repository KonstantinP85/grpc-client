package handler

import (
	"github.com/KonstantinP85/grpc-client/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		book := api.Group("/book")
		{
			book.GET("/", h.uploadNews)
			book.GET("/:id", h.getNews)
		}
	}

	return router
}
