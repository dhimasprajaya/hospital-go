package endpoint

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
)

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filename := filepath.Base(uuid.New().String())
	extension := filepath.Ext(file.Filename)
	if err := c.SaveUploadedFile(file, fmt.Sprintf("public/images/%s%s", filename, extension)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success Upload", "link": fmt.Sprintf("public/images/%s%s", filename, extension)})
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filename := filepath.Base(uuid.New().String())
	extension := filepath.Ext(file.Filename)
	if err := c.SaveUploadedFile(file, fmt.Sprintf("public/files/%s%s", filename, extension)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success Upload", "link": fmt.Sprintf("public/files/%s%s", filename, extension)})
}
