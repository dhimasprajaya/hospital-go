package endpoint

import (
	"andhiga.com/dhimasprajaya/go-vue-rs/config"
	"andhiga.com/dhimasprajaya/go-vue-rs/model"
	"andhiga.com/dhimasprajaya/go-vue-rs/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var list []model.User
	config.DB.Find(&list)
	c.JSON(http.StatusOK, list)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	var obj model.User

	// Record Not Found
	if config.DB.First(&obj, id).RecordNotFound() == true {
		c.JSON(http.StatusOK, gin.H{"message": "Record Not Found"})
		return
	}

	c.JSON(http.StatusOK, obj)
}

func CreateUser(c *gin.Context) {
	var json model.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password, err := util.HashPassword(json.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	json.Password = password

	if result := config.DB.Create(&json); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusCreated, json)
}

func UpdateUser(c *gin.Context) {
	var body model.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := body.Id
	var obj model.User
	// Record Not Found
	if config.DB.First(&obj, id).RecordNotFound() == true {
		c.JSON(http.StatusOK, gin.H{"message": "Record Not Found"})
		return
	}

	password, err := util.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	body.Password = password

	if result := config.DB.Save(&body); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Success Update", "data": body})
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var obj model.User

	// Record Not Found
	if config.DB.First(&obj, id).RecordNotFound() == true {
		c.JSON(http.StatusOK, gin.H{"message": "Record Not Found"})
		return
	}

	if err := config.DB.Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Success Delete", "data": obj})
	}
}
