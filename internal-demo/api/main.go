
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Define a route
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, world!",
        })
    })

    // Start serving the application
    router.Run()
}
