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
		auth.GET("/refresh", h.refreshToken)
	}
	api := router.Group("/api", h.userIdentity)
	{
		items := api.Group("/item")
		{
			items.POST("/", h.creatItem)
			items.GET("/", h.getAllItem)
			items.GET("/:item_id", h.getItemByID)
			items.DELETE(":item_id", h.deleteItem)
			//items.GET("/:title", h.getItemByTitle)
		}
		badHabits := api.Group("bad-habit")
		{
			badHabits.POST("/", h.createHabit)
			badHabits.GET("/", h.getAllHabits)
			badHabits.GET("/:habit_id", h.getHabitByID)
			badHabits.PATCH("/:habit_id", h.editEquivalentByID)
			badHabits.DELETE("/:habit_id", h.deleteHabit)
		}
		//mentalDev := api.Group("/mental")
		//{
		//	//mentalDev.GET("/", h.myInfo)
		//	mentalDev.POST("/", h.createCourse)
		//	mentalDev.GET("/:course_id", h.getCourseByID)
		//	mentalDev.DELETE("/:course_id", h.deleteCourseByID)
		//}
	}
	return router
}
