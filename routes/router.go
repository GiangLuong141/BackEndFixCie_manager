package routes

import (
	"example/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/employee", controllers.CreateEmployee())
	router.PUT("/employee/:id", controllers.UpdateEmployee())
	router.GET("/employee", controllers.GetListEmployee())
	router.GET("/employee/:id", controllers.GetDetailEmployee())
	router.DELETE("/employee/:id", controllers.DeleteEmployee())
	router.POST("/topic", controllers.CreateTopic())
	router.PUT("/topic/:id", controllers.UpdateTopic())
	router.GET("/topic", controllers.GetListTopic())
	router.GET("/topic/:id", controllers.GetDetailTopic())
	router.DELETE("topic/:id", controllers.DeleteEmployee())
}
