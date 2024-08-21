package handlers

import (
	"net/http"
	"strconv"
	"test/internal/model"
	"test/internal/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createUser(c *gin.Context) {
	var input model.User
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.services.User.Create(input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, utils.StatusResponse{Status: "user created"})
}

func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.User.GetAll()
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) getUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	user, err := h.services.User.GetById(userId)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) deleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	err = h.services.User.Delete(userId)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, utils.StatusResponse{Status: "user deleted"})
}
