package main

import (
	"log"
	"sample/config"
	"sample/router"
)

func main() {
	// Init DB
	config.InitDatabase()

	// Init Router
	r := router.InitRouter()
	log.Fatal(r.Run())

	//// Start Server Port :8000 Development
	//log.Fatal(r.Run(":8000"))
}
