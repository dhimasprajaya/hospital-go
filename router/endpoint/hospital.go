package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sample/config"
	"sample/model"
)

func GetHospitals(c *gin.Context) {
	var list []model.Hospital
	var obj model.Hospital

	// Handle QueryParam with Object Field Name
	if c.Bind(&obj) == nil {
		config.DB.Where(&obj).Find(&list)
		c.JSON(http.StatusOK, list)
		return
	}

	config.DB.Find(&list)
	c.JSON(http.StatusOK, list)
}

func GetHospitalById(c *gin.Context) {
	id := c.Param("id")
	var obj model.Hospital

	// Record Not Found
	if config.DB.First(&obj, id).RecordNotFound() == true {
		c.JSON(http.StatusOK, gin.H{"message": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, obj)
}

func CreateHospital(c *gin.Context) {
	var json model.Hospital
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

func UpdateHospital(c *gin.Context) {
	var body model.Hospital
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := body.ID
	var obj model.Hospital
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

func DeleteHospital(c *gin.Context) {
	id := c.Param("id")
	var obj model.Hospital

	// Record Not Found
	if config.DB.First(&obj, id).RecordNotFound() == true {
		c.JSON(http.StatusOK, gin.H{"message": "Record Not Found"})
		return
	}

	if err := config.DB.Where("id = ?", id).Delete(&model.Hospital{}).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Success Delete", "data": obj})
	}
}
