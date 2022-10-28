package handlers

import (
	"github.com/gin-gonic/gin"
	domain "github.com/ssharifzoda/levelup/internal/types"
	"strconv"
)

const (
	massage               = "Your operation completed successfully"
	positiveValidCategory = "He did not have this"
	negativeValidCategory = "You already have this"
)

type GetAllBadHabitsResponse struct {
	Data []domain.BadHabitOutput `json:"data"`
}

func (h *Handler) createHabit(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input domain.BadHabitInput
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, 400, "incorrect request")
		return
	}
	//call validate service method
	text, err := h.services.BadHabit.ValidateCategory(input.HabitCategoryId, userId)
	if text == negativeValidCategory || err != nil {
		c.JSON(400, text)
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
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	items, err := h.services.BadHabit.GetAll(userId)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
	}
	c.JSON(200, GetAllBadHabitsResponse{
		Data: items,
	})
}
func (h *Handler) getHabitByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("habit_id"))
	if err != nil {
		NewErrorResponse(c, 400, "invalid id param")
	}
	item, err := h.services.BadHabit.GetById(userId, id)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
	}
	c.JSON(200, item)
}
func (h *Handler) deleteHabit(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("habit_id"))
	if err != nil {
		NewErrorResponse(c, 400, "invalid id param")
	}
	text, err := h.services.BadHabit.DeleteHabitById(userId, id)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
	}
	c.JSON(200, text)
}
func (h *Handler) editEquivalentByID(c *gin.Context) {

}
