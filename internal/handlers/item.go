package handlers

import (
	"net/http"
	"strconv"
	"test/internal/model"
	"test/internal/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	var input model.Item
	if err := c.BindJSON(&input); err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.services.Item.Create(input)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, utils.StatusResponse{Status: "item created"})
}

func (h *Handler) getAllItems(c *gin.Context) {
	items, err := h.services.Item.GetAll()
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) getItemById(c *gin.Context) {
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	item, err := h.services.Item.GetById(itemId)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) deleteItem(c *gin.Context) {
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.NewErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	err = h.services.Item.Delete(itemId)
	if err != nil {
		utils.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, utils.StatusResponse{Status: "item deleted"})
}
