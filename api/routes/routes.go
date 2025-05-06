package routes

import "github.com/gin-gonic/gin"

// Shorten url
func ShortenURL(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "URL shortened successfully",
	})
}
