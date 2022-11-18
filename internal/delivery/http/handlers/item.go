package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssharifzoda/levelup/internal/types"
	"strconv"
)

func (h *Handler) creatItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input domain.Item
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, 400, "incorrect request")
		return
	}
	//call service method
	id, err := h.services.Diary.Create(userId, input)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
	}
	c.JSON(200, map[string]interface{}{
		"id": id,
	})
}

type getAllItemsResponse struct {
	Data []domain.Item `json:"data"`
}

func (h *Handler) getAllItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	pageNo, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		NewErrorResponse(c, 400, "invalid page params")
	}
	itemLimit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		NewErrorResponse(c, 400, "invalid page params")
	}
	items, err := h.services.Diary.GetAll(userId, pageNo, itemLimit)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
	}
	c.JSON(200, getAllItemsResponse{
		Data: items,
	})
}
func (h *Handler) getItemByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		NewErrorResponse(c, 400, "invalid id param")
	}
	item, err := h.services.Diary.GetById(userId, id)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
	}
	c.JSON(200, item)
}
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		NewErrorResponse(c, 400, "invalid id param")
	}
	text, err := h.services.Diary.DeleteItemById(userId, id)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
	}
	c.JSON(200, text)
}
func (h *Handler) getItemByTitle(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	title := c.Query("title")
	item, err := h.services.Diary.GetItemByTitle(userId, title)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
	}
	c.JSON(200, item)
}
