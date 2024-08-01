package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/nietthijmen/tavern/src/storage"
)

func CreateServer() {
	storage.InitialiseDriverMap()

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(AddIpMiddleware())

	r.GET("/up", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "up",
		})
	})

	r.GET("/public/:slug", downloadFile)

	//generate a trusted group
	group := r.Group("/trusted")
	group.Use(VerifyKeyMiddleware())
	group.POST("/upload", uploadFile)

	r.Run()
}
