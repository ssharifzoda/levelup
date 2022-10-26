package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
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
	//parse token
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, 401, "invalid token, pls take again")
		return
	}
	c.Set(userCtx, userId)
}
func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		NewErrorResponse(c, 500, "invalid operation user id not found")
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		NewErrorResponse(c, 500, "user id is of invalid type")
		return 0, errors.New("user id not found")
	}
	return idInt, nil
}
