package handlers

import (
	"github.com/gin-gonic/gin"
	domain "github.com/ssharifzoda/levelup/internal/types"
	"strconv"
)

func (h *Handler) createCourse(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input domain.CourseInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, 400, "incorrect request")
		return
	}
	//call service method
	id, err := h.services.MentalDevelopment.Create(userId, input)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
	}
	c.JSON(200, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) getCourseByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		NewErrorResponse(c, 400, "invalid id param")
	}
	item, err := h.services.MentalDevelopment.GetById(userId, id)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
		return
	}
	c.Writer.Header().Set("Content-Type", "multiparty-data")
}
func (h *Handler) deleteCourseByID(c *gin.Context) {

}
