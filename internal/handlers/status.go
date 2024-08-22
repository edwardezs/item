package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	resputils "test/internal/utils"
)

func (h *Handler) getAPIStatus(ctx *gin.Context) {
	status, err := h.services.Status.GetStatus()
	if err != nil {
		resputils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	resputils.SuccessResponse(ctx, status)
}

func (h *Handler) changeAPIStatus(ctx *gin.Context) {
	status, err := strconv.ParseBool(ctx.Param("read_only"))
	if err != nil {
		resputils.ErrorResponse(ctx, http.StatusBadRequest, "invalid api status param")
		return
	}

	newStatus, err := h.services.Status.ChangeStatus(status)
	if err != nil {
		resputils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	resputils.SuccessResponse(ctx, newStatus)
}
