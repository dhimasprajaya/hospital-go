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

	// Start Server
	log.Fatal(r.Run(":8080"))

	// Start Server Port :8000 Development
	//log.Fatal(r.Run(":8000"))
}
