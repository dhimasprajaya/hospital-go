package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"hospital-go/model"
)

var DB *gorm.DB

func InitDatabase() {
	// Configure DB Connection
	var err error
	DB, err = gorm.Open(Config.DbType, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		Config.DbHost, Config.DbPort, Config.DbUser, Config.DbName, Config.DbPassword))
	if err != nil {
		panic("Failed to connect database")
	}
	//defer DB.Close()

	// Migrate the schema
	DB.AutoMigrate(
		&model.User{},
		&model.Hospital{},
		&model.Doctor{},
		&model.Patient{},
		&model.MedicalHistory{},
	)
}
