package handlers

import (
	"items/internal/service"

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
		items := api.Group("/items")
		{
			items.GET("/", h.getAllItems)
			items.GET("/:id", h.getItemById)
			items.POST("/:id", h.createItem)
			items.DELETE("/:id", h.deleteItem)
		}

		status := api.Group("/status")
		{
			status.GET("/", h.getAPIStatus)
		}
	}

	return router
}
