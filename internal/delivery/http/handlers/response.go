package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssharifzoda/levelup/pkg/logging"
)

type Error struct {
	Massage string `json:"massage"`
}

func NewErrorResponse(c *gin.Context, statusCode int, massage string) {
	logger := logging.GetLogger()
	logger.Error(massage)
	c.AbortWithStatusJSON(statusCode, Error{massage})
}
