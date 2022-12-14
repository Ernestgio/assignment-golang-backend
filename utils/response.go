package utils

import "github.com/gin-gonic/gin"

func ResponseWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}
