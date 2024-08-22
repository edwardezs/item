package handlers

import (
	"fmt"
	"net/http"
	"test/internal/model"
	"test/internal/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createUser(ctx *gin.Context) {
	var input model.User
	if err := ctx.BindJSON(&input); err != nil {
		utils.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.services.User.Create(input)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(ctx, fmt.Sprintf("user %s created", input.Name))
}

func (h *Handler) getAllUsers(ctx *gin.Context) {
	users, err := h.services.User.GetAll()
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(ctx, users)
}

func (h *Handler) getUserById(ctx *gin.Context) {
	userId, err := utils.ParseId(ctx)
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, "invalid user id param")
		return
	}

	user, err := h.services.User.GetById(userId)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(ctx, user)
}

func (h *Handler) deleteUser(ctx *gin.Context) {
	userId, err := utils.ParseId(ctx)
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, "invalid user id param")
		return
	}

	name, err := h.services.User.Delete(userId)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(ctx, fmt.Sprintf("user %s deleted", name))
}
