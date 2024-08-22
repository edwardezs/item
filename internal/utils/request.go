package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	id       = "id"
	readOnly = "read_only"
)

func ParseId(ctx *gin.Context) (int, error) {
	return strconv.Atoi(ParseParam(ctx, id))
}

func ParceReadOnly(ctx *gin.Context) (bool, error) {
	return strconv.ParseBool(ParseParam(ctx, readOnly))
}

func ParseParam(ctx *gin.Context, param string) string {
	return ctx.Param(param)
}
