package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	domain "github.com/ssharifzoda/levelup/internal/types"
	"os"
	"strconv"
)

func (h *Handler) createBodyCourse(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input domain.BodyCourseInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, 400, "incorrect request")
		return
	}
	//call validate service
	text, err := h.services.PhysicianDevelopment.ValidateCategory(input.TrainCategoryId, userId)
	if text == negativeValidCategory || err != nil {
		c.JSON(400, text)
		return
	}
	//call service method
	id, err := h.services.PhysicianDevelopment.Create(userId, input)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
	}
	c.JSON(200, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) getBodyCourseByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		NewErrorResponse(c, 400, "invalid id param")
	}
	item, err := h.services.PhysicianDevelopment.GetById(userId, id)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
		return
	}
	c.JSON(200, item)
}
func (h *Handler) getVideoByCourse(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		NewErrorResponse(c, 400, "invalid id param")
		return
	}
	item, err := h.services.PhysicianDevelopment.GetById(userId, id)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
		return
	}
	video, err := os.ReadFile(item.Video)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
		return
	}
	if _, err = c.Writer.Write(video); err != nil {
		logrus.Print(err)
		return
	}
}
func (h *Handler) getPlaylist(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		NewErrorResponse(c, 400, "invalid id param")
		return
	}
	item, err := h.services.PhysicianDevelopment.GetById(userId, id)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
		return
	}
	c.File(item.Playlist)
}
func (h *Handler) diet(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		NewErrorResponse(c, 400, "invalid id param")
		return
	}
	item, err := h.services.PhysicianDevelopment.GetById(userId, id)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
		return
	}
	diet, err := os.ReadFile(item.Diet)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
		return
	}
	if _, err = c.Writer.Write(diet); err != nil {
		logrus.Print(err)
		return
	}
}

func (h *Handler) trainPlan(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		NewErrorResponse(c, 400, "invalid id param")
		return
	}
	item, err := h.services.PhysicianDevelopment.GetById(userId, id)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
		return
	}
	trainPlan, err := os.ReadFile(item.TrainPlan)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
		return
	}
	if _, err = c.Writer.Write(trainPlan); err != nil {
		logrus.Print(err)
		return
	}
}

func (h *Handler) deleteBodyCourse(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		NewErrorResponse(c, 400, "invalid id param")
		return
	}
	text, err := h.services.PhysicianDevelopment.DeleteCourseById(userId, id)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
		return
	}
	c.JSON(200, text)
}
