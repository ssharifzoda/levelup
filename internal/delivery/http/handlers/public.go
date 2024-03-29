package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	domain "github.com/ssharifzoda/levelup/internal/types"
	"os"
)

const (
	public = "public.txt"
	test   = "test.txt"
)

func (h *Handler) getPublic(c *gin.Context) {
	public, err := os.ReadFile(viper.GetString("storage.public") + public)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
		return
	}
	if _, err = c.Writer.Write(public); err != nil {
		logrus.Print(err)
		return
	}
}
func (h *Handler) receivePublic(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input domain.Public
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, 400, "incorrect request")
		return
	}
	err = h.services.Public.ReceivePublic(userId, input)
	if err != nil {
		NewErrorResponse(c, 500, "internal error")
		return
	}
	c.JSON(200, massage)
}
func (h *Handler) getTest(c *gin.Context) {
	test, err := os.ReadFile(viper.GetString("storage.public") + test)
	if err != nil {
		NewErrorResponse(c, 500, err.Error())
		return
	}
	if _, err = c.Writer.Write(test); err != nil {
		logrus.Print(err)
		return
	}
}
func (h *Handler) doTest(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input domain.Test
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, 400, "incorrect request")
		return
	}
	temperament, err := h.services.Public.DoTest(userId, input)
	if err != nil {
		NewErrorResponse(c, 500, "internal error")
		return
	}
	c.JSON(200, temperament)
}
