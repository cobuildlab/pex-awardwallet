package handlers

import (
	"github.com/gin-gonic/gin"
)

//LaunchResponseJSON Launches a JSON message to the client
func LaunchResponseJSON(c *gin.Context, code int, response gin.H) {
	c.JSON(code, gin.H{
		"statusCode": code,
		"response":   response,
	})
}

//LaunchResponseErrorJSON Throws a JSON error to the client
func LaunchResponseErrorJSON(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{
		"statusCode": code,
		"response": gin.H{
			"error": err.Error(),
		},
	})
}
