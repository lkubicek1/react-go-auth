package routes

import (
	"api/controllers"
	"github.com/gin-gonic/gin"
)

func CounterRoutes(router *gin.Engine) {
	counter := router.Group("/counter")
	{
		counter.GET("/current", controllers.GetCurrentCounter)
		counter.POST("/increment", controllers.IncrementCounter)
		counter.POST("/reset", controllers.ResetCounter)
	}
}
