package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sample/config"
	"sample/model"
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
		c.JSON(http.StatusNoContent, gin.H{"message": "Record Not Found"})
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

	if result := config.DB.Create(&json); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusCreated, json)
}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
