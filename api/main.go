package main

import (
	"api/config" // Assuming this package contains your EnvVars and other configurations
	"api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize your configuration
	// Assuming you have a function or method to load your environment variables
	envConfig := config.LoadEnv()

	// Now pass this configuration to your AdminRoutes
	routes.AdminRoutes(router, envConfig)

	// Start the server
	router.Run(":8080")
}
