package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ssharifzoda/levelup/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		items := api.Group("/item")
		{
			items.POST("/", h.creatItem)
			items.GET("/", h.getAllItem)
			items.GET("/:item_id", h.getItemByID)
			items.DELETE(":item_id", h.deleteItem)
		}
	}
	return router
}
