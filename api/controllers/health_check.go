package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	//return a JSON response
	c.JSON(http.StatusOK, gin.H{
		"message": "Healthy!",
	})
}
