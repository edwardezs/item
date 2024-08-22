package handlers

import (
	"fmt"
	"net/http"

	"test/internal/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createUser(ctx *gin.Context) {
	var input model.User
	if err := ctx.BindJSON(&input); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.services.User.Create(input)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(ctx, fmt.Sprintf("user %s created", input.Name))
}

func (h *Handler) getAllUsers(ctx *gin.Context) {
	users, err := h.services.User.GetAll()
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(ctx, users)
}

func (h *Handler) getUserById(ctx *gin.Context) {
	userId, err := parseId(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid user id param")
		return
	}

	user, err := h.services.User.GetById(userId)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(ctx, user)
}

func (h *Handler) deleteUser(ctx *gin.Context) {
	userId, err := parseId(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid user id param")
		return
	}

	name, err := h.services.User.Delete(userId)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(ctx, fmt.Sprintf("user %s deleted", name))
}
