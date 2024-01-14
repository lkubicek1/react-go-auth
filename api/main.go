package main

import (
	"api/middleware"
	"api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	router := gin.Default()

	// Initialize environment configuration
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowWildcard:    true,
	}))

	// Allow all origins for CORS - For testing only
	//router.Use(cors.Default())

	// Routes
	routes.StatusRoutes(router)
	routes.CounterRoutes(router)

	// Secured Routes
	auth := middleware.NewAuthMiddleware()
	routes.UserCounterRoutes(router, auth)

	// Start the server
	err = router.Run(":8080")
	if err != nil {
		return
	}
}
