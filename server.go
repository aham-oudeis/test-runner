package main

import (
	"test-runner/api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	router.GET("/health", controllers.HealthCheck)

	router.POST("/assertions/validate", controllers.ValidateAssertions)

	router.POST("/requests/execute", controllers.ExecuteRequests)

	// Run the server on port 8080
	router.Run(":8080")
}