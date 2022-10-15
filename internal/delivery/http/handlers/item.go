package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssharifzoda/levelup/internal/domain"
)

func (h *Handler) creatItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input domain.Item
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, 403, "incorrect request")
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
func (h *Handler) getAllItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	items, err := h.services.Diary.GetAll(userId)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
	}

}
func (h *Handler) getItemByID(c *gin.Context) {

}
func (h *Handler) deleteItem(c *gin.Context) {

}
