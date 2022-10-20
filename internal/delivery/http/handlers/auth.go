package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/ssharifzoda/levelup/internal/types"
	"net/http"
	"time"
)

const Validate = "You are already registered"

func (h *Handler) signUp(c *gin.Context) {
	var input domain.User
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	massage, err := h.services.Authorization.Validate(input.Username, input.Password)
	if massage == Validate && err == nil {
		c.JSON(400, massage)
		return
	}
	if err != nil {
		logrus.Println(err)
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		logrus.Println(err)
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, map[string]interface{}{
		"active until": time.Now().Add(time.Minute * 10).Format(time.Kitchen),
		"your token":   token,
	})
}
