package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	domain "github.com/ssharifzoda/levelup/internal/types"
	"os"
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
	//call validate service
	text, err := h.services.MentalDevelopment.ValidateCategory(input.MentalCategoryId, userId)
	if text == negativeValidCategory || err != nil {
		c.JSON(400, text)
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
	c.JSON(200, item)
}

func (h *Handler) getAudio(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		NewErrorResponse(c, 400, "invalid id param")
		return
	}
	item, err := h.services.MentalDevelopment.GetById(userId, id)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
		return
	}
	audio, err := os.ReadFile(item.Audio)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
		return
	}
	if _, err = c.Writer.Write(audio); err != nil {
		logrus.Print(err)
		return
	}
}

func (h *Handler) getBook(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		NewErrorResponse(c, 400, "invalid id param")
		return
	}
	item, err := h.services.MentalDevelopment.GetById(userId, id)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
		return
	}
	book, err := os.ReadFile(item.Book)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
		return
	}
	if _, err = c.Writer.Write(book); err != nil {
		logrus.Print(err)
		return
	}
}

func (h *Handler) deleteCourseByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		NewErrorResponse(c, 400, "invalid id param")
		return
	}
	text, err := h.services.MentalDevelopment.DeleteCourseById(userId, id)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
		return
	}
	c.JSON(200, text)
}
func (h *Handler) GetMentalCategories(c *gin.Context) {
	categories, err := h.services.MentalDevelopment.GetCategories()
	if err != nil {
		NewErrorResponse(c, 500, "internal error")
	}
	c.JSON(200, categories)
}
