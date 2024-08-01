package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/nietthijmen/tavern/src/database"
)

func VerifyKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(401, gin.H{
				"error": "Authorization header is missing",
			})
			c.Abort()
			return
		}

		_, err := database.GetApiKey(header)
		if err != nil {
			c.JSON(401, gin.H{
				"error": "Invalid API key",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
func AddIpMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("CF-Connecting-IP") != "" {
			c.Set("ip", c.GetHeader("CF-Connecting-IP"))
			c.Next()
			return
		}

		ip := c.ClientIP()
		c.Set("ip", ip)
		c.Next()
	}
}
