package main

import (
	"api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.AdminRoutes(router)
	router.Run(":8080")
}
