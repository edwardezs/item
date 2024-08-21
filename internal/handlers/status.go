package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAPIStatus(c *gin.Context) {
	status, err := h.services.Status.GetStatus()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, status)
}
