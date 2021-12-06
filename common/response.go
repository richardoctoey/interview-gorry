package common

import "github.com/gin-gonic/gin"

func Error(c *gin.Context, err error) {
	c.JSON(200, gin.H{
		"success": false,
		"message": err.Error(),
	})
}

func OK(c *gin.Context, message string, data interface{}) {
	if data != nil {
		c.JSON(200, gin.H{
			"success": true,
			"message": message,
			"data": data,
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"message": message,
	})
}
