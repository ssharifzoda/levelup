package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/ssharifzoda/levelup/internal/types"
	"net/http"
	"strings"
	"unicode"
)

const (
	errLength         = "password < 9 or > 25"
	negativeValidUser = "user already registered"
	newUser           = "user not registered in database"
	noCyr             = "dont typing Cyrillic"
)

func (h *Handler) signUp(c *gin.Context) {
	var input domain.User
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := passwordValidate(input.Password)
	if err != nil {
		c.JSON(400, err)
		return
	}
	massage, err := h.services.Authorization.UserValidate(input.Username, input.Password)
	if massage != newUser {
		c.JSON(400, negativeValidUser)
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
	token, refreshToken, err := h.services.Authorization.GenerateTokens(input.Username, input.Password)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, map[string]interface{}{
		"token":         token,
		"refresh-token": refreshToken,
	})
}
func (h *Handler) refreshToken(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		NewErrorResponse(c, 401, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(c, 401, "invalid auth header")
		return
	}
	username, passwordHash, err := h.services.Authorization.ParseRefreshToken(headerParts[1])
	if err != nil {
		return
	}
	token, refreshToken, err := h.services.Authorization.GenerateTokens(username, passwordHash)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, map[string]interface{}{
		"token":         token,
		"refresh-token": refreshToken,
	})
}
func passwordValidate(password string) error {
	if len(password) < 9 && len(password) > 25 {
		return errors.New(errLength)
	}
	pass := []rune(password)
	for _, val := range pass {
		if unicode.Is(unicode.Cyrillic, val) == true {
			return errors.New(noCyr)
		}
	}
	return nil
}
