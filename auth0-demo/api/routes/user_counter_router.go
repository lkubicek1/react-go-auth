package routes

import (
	"api/controllers"
	"api/middleware"
	"github.com/gin-gonic/gin"
)

// UserCounterRoutes sets up the routes for user-specific counters.
// It takes the main router and a middleware function as arguments.
func UserCounterRoutes(router *gin.Engine, auth *middleware.AuthMiddleware) {
	// Create a new group for user-specific counters
	userCounter := router.Group("/user-counter")

	// Apply the middleware to this group
	userCounter.Use(auth.ValidateToken)

	// Define routes within this group
	userCounter.GET("/current", controllers.GetUserCounter)
	userCounter.POST("/increment", controllers.IncrementUserCounter)
	userCounter.POST("/reset", controllers.ResetUserCounter)
}
