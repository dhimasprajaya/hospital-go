package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"hospital-go/util"
	"net/http"
)

type Header struct {
	Authorization string `json:"authorization"`
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := Header{}
		if err := c.ShouldBindHeader(&header); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		token := header.Authorization
		var msg = "OK"

		if token == "" {
			msg = "Token Invalid"
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					msg = "Token Expired"
				default:
					msg = "Token Error"
				}
			}
		}

		if msg != "OK" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": msg})
			c.Abort()
			return
		}

		c.Next()
	}
}
