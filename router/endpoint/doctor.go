package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sample/config"
	"sample/model"
)

func GetDoctors(c *gin.Context) {
	var list []model.Doctor
	config.DB.Find(&list)
	c.JSON(http.StatusOK, list)
}

func GetDoctorById(c *gin.Context) {
	id := c.Param("id")
	var obj model.Doctor

	// Record Not Found
	if config.DB.First(&obj, id).RecordNotFound() == true {
		c.JSON(http.StatusOK, gin.H{"message": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, obj)
}

func CreateDoctor(c *gin.Context) {
	var json model.Doctor
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := config.DB.Create(&json); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusCreated, json)
}

func UpdateDoctor(c *gin.Context) {
	var body model.Doctor
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := body.ID
	var obj model.Doctor
	// Record Not Found
	if config.DB.First(&obj, id).RecordNotFound() == true {
		c.JSON(http.StatusOK, gin.H{"message": "Record Not Found"})
		return
	}

	if result := config.DB.Save(&body); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Success Update", "data": body})
	}
}

func DeleteDoctor(c *gin.Context) {
	id := c.Param("id")
	var obj model.Doctor

	// Record Not Found
	if config.DB.First(&obj, id).RecordNotFound() == true {
		c.JSON(http.StatusOK, gin.H{"message": "Record Not Found"})
		return
	}

	if err := config.DB.Where("id = ?", id).Delete(&model.Doctor{}).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Success Delete", "data": obj})
	}
}
