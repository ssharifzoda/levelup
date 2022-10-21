package handlers

import (
	"github.com/gin-gonic/gin"
	domain "github.com/ssharifzoda/levelup/internal/types"
	"strconv"
)

const massage = "Your operation completed successfully"

type GetAllBadHabitsResponse struct {
	Data []domain.BadHabit `json:"data"`
}

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
func (h *Handler) doExercise(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("habit_id"))
	if err != nil {
		NewErrorResponse(c, 400, "invalid id param")
	}
	var input domain.DoExercise
	text, err := h.services.BadHabit.DoExercise(userId, id, input)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
	}
	c.JSON(200, text)
}
