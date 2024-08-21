package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"test/internal/utils"
)

func (h *Handler) getAPIStatus(c *gin.Context) {
	status, err := h.services.Status.GetStatus()
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, status)
}
