package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"sample/router/endpoint"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	// Main Router
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"_code":             200,
			"_active":           true,
			"_timestamp":        time.Now(),
			"_message":          "REST API for RS Website",
			"_base_url":         "https://go-vue-rs.herokuapp.com/",
			"endpoint_user":     "api/user",
			"endpoint_hospital": "api/hospital",
			"endpoint_doctor":   "api/doctor",
			"endpoint_patient":  "api/patient",
		})
	})

	// API Router
	api := r.Group("/api")
	{
		user := api.Group("/user")
		{
			user.GET("/", endpoint.GetUsers)
			user.GET("/:id", endpoint.GetUserById)
			user.POST("/", endpoint.CreateUser)
			user.PUT("/", endpoint.UpdateUser)
			user.DELETE("/:id", endpoint.DeleteUser)
		}
	}

	return r
}
