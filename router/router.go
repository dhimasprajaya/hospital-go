package router

import (
	"andhiga.com/dhimasprajaya/go-vue-rs/middleware"
	"andhiga.com/dhimasprajaya/go-vue-rs/router/endpoint"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// Enable CORS
	r.Use(cors.Default())
	// Templates HTML
	r.LoadHTMLGlob("templates/*")
	// Static Directory
	r.Static("/public", "./public")
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Main Router
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Auth
	r.POST("/login", endpoint.Login)

	// API Router
	api := r.Group("/api")
	api.Use(middleware.JWT())
	{
		user := api.Group("/user")
		{
			user.GET("/", endpoint.GetUsers)
			user.GET("/:id", endpoint.GetUserById)
			user.POST("/", endpoint.CreateUser)
			user.PUT("/", endpoint.UpdateUser)
			user.DELETE("/:id", endpoint.DeleteUser)
		}
		hospital := api.Group("/hospital")
		{
			hospital.GET("/", endpoint.GetHospitals)
			hospital.GET("/:id", endpoint.GetHospitalById)
			hospital.POST("/", endpoint.CreateHospital)
			hospital.PUT("/", endpoint.UpdateHospital)
			hospital.DELETE("/:id", endpoint.DeleteHospital)
		}
		doctor := api.Group("/doctor")
		{
			doctor.GET("/", endpoint.GetDoctors)
			doctor.GET("/:id", endpoint.GetDoctorById)
			doctor.POST("/", endpoint.CreateDoctor)
			doctor.PUT("/", endpoint.UpdateDoctor)
			doctor.DELETE("/:id", endpoint.DeleteDoctor)
		}
		patient := api.Group("/patient")
		{
			patient.GET("/", endpoint.GetPatients)
			patient.GET("/:id", endpoint.GetPatientById)
			patient.POST("/", endpoint.CreatePatient)
			patient.PUT("/", endpoint.UpdatePatient)
			patient.DELETE("/:id", endpoint.DeletePatient)
		}
		upload := api.Group("/upload")
		{
			upload.POST("/image", endpoint.UploadImage)
			upload.POST("/file", endpoint.UploadFile)
		}
	}

	return r
}
