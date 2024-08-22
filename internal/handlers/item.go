package handlers

import (
	"fmt"
	"net/http"

	"test/internal/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(ctx *gin.Context) {
	var input model.Item
	if err := ctx.BindJSON(&input); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.services.Item.Create(input)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(ctx, fmt.Sprintf("item %s created", input.Title))
}

func (h *Handler) getAllItems(ctx *gin.Context) {
	items, err := h.services.Item.GetAll()
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(ctx, items)
}

func (h *Handler) getItemById(ctx *gin.Context) {
	itemId, err := parseId(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid item id param")
		return
	}

	item, err := h.services.Item.GetById(itemId)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(ctx, item)
}

func (h *Handler) deleteItem(ctx *gin.Context) {
	itemId, err := parseId(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid item id param")
		return
	}

	title, err := h.services.Item.Delete(itemId)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	successResponse(ctx, fmt.Sprintf("item with %s deleted", title))
}
