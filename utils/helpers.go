package utils

import (
	// "net/http"
	"github.com/gin-gonic/gin" // Ensure this import is included
)

// Custom response helper
func SendResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}
