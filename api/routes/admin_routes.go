package routes

import (
	"api/auth"
	"api/config"
	"api/controllers"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine, authConfig config.EnvVars) {
	// Initialize the auth middleware
	authMiddleware := auth.NewMiddleware(authConfig)

	admin := router.Group("/admin")
	// Apply the middleware
	admin.Use(authMiddleware.ValidateToken)
	{
		admin.GET("/status", controllers.GetStatus)
	}
}
