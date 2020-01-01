package controllers

import "github.com/gin-gonic/gin"

func GetHome() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"system":  "Is up is up!",
			"message": "Everything is up in your life?",
		})
	}
}
