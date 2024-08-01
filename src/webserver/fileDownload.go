package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/nietthijmen/tavern/src/database"
	"github.com/nietthijmen/tavern/src/storage"
)

func downloadFile(c *gin.Context) {
	slug := c.Param("slug")

	bySlug, err := database.FindObjectBySlug(slug)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "File not found",
		})
		return
	}

	driver := storage.GetInitialisedDriver(bySlug.Bucket)

	reader, err := driver.StreamFile(bySlug.FilePath)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error streaming file",
		})
		return
	}

	defer reader.Close()

	c.Header("Content-Type", "application/octet-stream")
	for {
		buf := make([]byte, 1024)
		n, err := reader.Read(buf)
		if err != nil {
			break
		}

		c.Writer.Write(buf[:n])
	}
}
