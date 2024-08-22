package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"test/internal/utils"
)

func (h *Handler) getAPIStatus(ctx *gin.Context) {
	status, err := h.services.Status.GetStatus()
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(ctx, status)
}

func (h *Handler) changeAPIStatus(ctx *gin.Context) {
	status, err := utils.ParceReadOnly(ctx)
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, "invalid api status param")
		return
	}

	newStatus, err := h.services.Status.ChangeStatus(status)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(ctx, fmt.Sprintf("api read-only mode is %t", newStatus))
}
