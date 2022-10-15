package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) creatDiary(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(200, map[string]interface{}{
		"id": id,
	})

}
func (h *Handler) getAllDiary(c *gin.Context) {

}
func (h *Handler) getDiaryByID(c *gin.Context) {

}
func (h *Handler) deleteDiary(c *gin.Context) {

}
