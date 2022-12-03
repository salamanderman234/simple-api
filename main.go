package main

import (
	"github.com/joho/godotenv"
	"github.com/salamanderman234/simple-api/database"
	"github.com/salamanderman234/simple-api/routes"
)

func main() {
	// setup dotenv
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// setup database
	err = database.ConnectAndMigrate()
	if err != nil {
		panic(err)
	}

	// setup router
	router := routes.Init()
	router.Logger.Fatal(router.Start(":3000"))
}
