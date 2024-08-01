package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nietthijmen/tavern/src/database"
	"github.com/nietthijmen/tavern/src/storage"
	"github.com/nietthijmen/tavern/src/utils"
	"github.com/rs/zerolog/log"
)

func uploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"error": "File is missing",
		})
		return
	}

	bucket, err := database.GetAvailableBucket(int(file.Size))
	if err != nil {
		log.Error().Err(err).Msg("Error getting available bucket")
		c.JSON(500, gin.H{
			"error": "No available bucket",
		})
		return
	}

	reader, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error opening file",
		})
		return
	}

	defer reader.Close()
	localFile := storage.LocalFile{
		Name:   file.Filename,
		Size:   int(file.Size),
		Reader: reader,
	}

	driver := storage.GetInitialisedDriver(bucket)
	remoteFile, err := driver.UploadFile(localFile)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error uploading file",
		})
		return
	}

	object := database.StorageObject{
		Slug:       utils.Slugify(bucket.Name + "_" + uuid.New().String() + "_" + file.Filename),
		Size:       remoteFile.Size,
		FileType:   file.Header.Get("Content-Type"),
		FilePath:   remoteFile.Path,
		UploadedBy: c.GetString("ip"),
		Bucket:     bucket,
	}

	err = database.CreateObject(object)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error creating object",
		})
		return
	}

	c.JSON(200, gin.H{
		"slug":      object.Slug,
		"size":      object.Size,
		"file_type": object.FileType,
	})
}
