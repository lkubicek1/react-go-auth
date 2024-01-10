package routes

import (
	"api/controllers"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine) {
	admin := router.Group("/admin")
	{
		admin.GET("/status", controllers.GetStatus)
	}
}
