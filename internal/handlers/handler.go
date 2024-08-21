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
			items.POST("/", h.createItem)
			items.GET("/:id", h.getItemById)
			items.DELETE("/:id", h.deleteItem)
		}

		user := api.Group("/users")
		{
			user.GET("/", h.getAllUsers)
			user.POST("/", h.createUser)
			user.GET("/:id", h.getUserById)
			user.DELETE("/:id", h.deleteUser)
		}

		status := api.Group("/status")
		{
			status.GET("/", h.getAPIStatus)
			// status.GET("/", h.getTableStatus)
			// status.POST("/:table_name", h.changeAPIStatus)
			// status.PUT("/:table_name", h.changeTableStatus)
		}
	}

	return router
}
