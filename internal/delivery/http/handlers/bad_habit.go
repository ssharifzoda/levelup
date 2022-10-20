package handlers

import (
	"github.com/gin-gonic/gin"
	domain "github.com/ssharifzoda/levelup/internal/types"
)

const massage = "Your operation completed successfully"

func (h *Handler) createHabit(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input domain.BadHabit
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, 400, "incorrect request")
		return
	}
	//call service method
	id, err := h.services.BadHabit.Create(userId, input)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
	}
	c.JSON(200, map[string]interface{}{
		"massage": massage,
		"id":      id,
	})
}
func (h *Handler) getAllHabits(c *gin.Context) {

}
func (h *Handler) getHabitByID(c *gin.Context) {

}
func (h *Handler) deleteHabit(c *gin.Context) {

}
func (h *Handler) doExercise(c *gin.Context) {

}
