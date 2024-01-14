package routes

import (
	"api/controllers"
	"github.com/gin-gonic/gin"
)

func StatusRoutes(router *gin.Engine) {
	status := router.Group("/status")
	{
		status.GET("/", controllers.GetStatus)
	}
}
