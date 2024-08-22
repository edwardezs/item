package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAPIStatus(ctx *gin.Context) {
	status, err := h.services.Status.GetStatus()
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(ctx, status)
}

func (h *Handler) changeAPIStatus(ctx *gin.Context) {
	status, err := parceReadOnly(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid api status param")
		return
	}

	newStatus, err := h.services.Status.ChangeStatus(status)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(ctx, fmt.Sprintf("api read-only mode is %t", newStatus))
}
