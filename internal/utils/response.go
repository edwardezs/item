package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type response struct {
	Status any `json:"status"`
	Code   int `json:"code"`
}

func ErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, response{Status: message, Code: statusCode})
}

func SuccessResponse(ctx *gin.Context, status any) {
	logrus.Info(status)
	ctx.JSON(http.StatusOK, response{Status: status, Code: http.StatusOK})
}
