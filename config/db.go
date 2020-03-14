package config

import (
	"andhiga.com/dhimasprajaya/go-vue-rs/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDatabase() {
	// Configure DB Connection
	var err error
	DB, err = gorm.Open("postgres", "host=ec2-34-200-116-132.compute-1.amazonaws.com port=5432 user=qlqzzfolomnpzb dbname=dbsl1nu7bebsi2 password=ab36db424e94b4cb2d5626ce3499d6344838b897f45250642fdce76c39e3dfa7")
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
