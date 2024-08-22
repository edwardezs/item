package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type response struct {
	Response   any `json:"response"`
	StatusCode int `json:"status_code"`
}

func Error(ctx *gin.Context, statusCode int, message any) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, response{Response: message, StatusCode: statusCode})
}

func Success(ctx *gin.Context, message any) {
	logrus.Info(message)
	ctx.JSON(http.StatusOK, response{Response: message, StatusCode: http.StatusOK})
}
