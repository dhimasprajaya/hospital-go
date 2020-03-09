package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// Configure DB
	db, err := gorm.Open("postgres", "host=ec2-34-200-116-132.compute-1.amazonaws.com port=5432 user=qlqzzfolomnpzb dbname=dbsl1nu7bebsi2 password=ab36db424e94b4cb2d5626ce3499d6344838b897f45250642fdce76c39e3dfa7")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Setup Route

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"_code":             200,
			"_active":           true,
			"_timestamp":        time.Now(),
			"_message":          "REST API for RS Website",
			"_base_url":         "https://go-vue-rs.herokuapp.com/",
			"endpoint_hospital": "api/hospital",
			"endpoint_doctor":   "api/doctor",
			"endpoint_patient":  "api/patient",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":      "Dhimas Prajaya",
			"gender":    "Male",
			"birthday":  "28/07/1990",
			"is_active": true,
			"age":       29,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//_ = r.Run("0.0.0.0:8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
