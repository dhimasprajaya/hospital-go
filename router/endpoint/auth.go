package endpoint

import (
	"andhiga.com/dhimasprajaya/go-vue-rs/config"
	"andhiga.com/dhimasprajaya/go-vue-rs/model"
	"andhiga.com/dhimasprajaya/go-vue-rs/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var json Auth
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Record Not Found
	var user model.User
	if config.DB.First(&user, "email = ? ", json.Email).RecordNotFound() == true {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Email Not Found"})
		return
	}

	isMatch := util.CheckPasswordHash(json.Password, user.Password)
	if isMatch != true {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong Password"})
		return
	}

	token, err := util.GenerateToken(json.Email, json.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
