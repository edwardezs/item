package handlers

import (
	"net/http"
	"strconv"

	"test/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	id       = "id"
	readOnly = "read_only"
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
			items.GET("", h.getAllItems)
			items.POST("", h.createItem)
			items.GET("/:id", h.getItemById)
			items.DELETE("/:id", h.deleteItem)
		}

		user := api.Group("/users")
		{
			user.GET("", h.getAllUsers)
			user.POST("", h.createUser)
			user.GET("/:id", h.getUserById)
			user.DELETE("/:id", h.deleteUser)
		}

		status := api.Group("/status")
		{
			status.GET("", h.getAPIStatus)
			status.POST("/:read_only", h.changeAPIStatus)
		}
	}

	return router
}

func parseParam(ctx *gin.Context, param string) string {
	return ctx.Param(param)
}

func parseId(ctx *gin.Context) (int, error) {
	return strconv.Atoi(parseParam(ctx, id))
}

func parceReadOnly(ctx *gin.Context) (bool, error) {
	return strconv.ParseBool(parseParam(ctx, readOnly))
}

type response struct {
	Response   any `json:"response"`
	StatusCode int `json:"status_code"`
}

func errorResponse(ctx *gin.Context, statusCode int, message any) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, response{Response: message, StatusCode: statusCode})
}

func successResponse(ctx *gin.Context, message any) {
	logrus.Info(message)
	ctx.JSON(http.StatusOK, response{Response: message, StatusCode: http.StatusOK})
}
