package router

import (
	"github.com/gin-gonic/gin"
	"sample/router/endpoint"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// Main Router
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"_code":             200,
			"_active":           true,
			"_timestamp":        time.Now(),
			"_message":          "REST API for RS Website",
			"_base_url":         "https://go-vue-rs.herokuapp.com/",
			"endpoint_hospital": "endpoint/hospital",
			"endpoint_doctor":   "endpoint/doctor",
			"endpoint_patient":  "endpoint/patient",
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
			user.DELETE("/", endpoint.DeleteUser)
		}
	}

	return r
}
