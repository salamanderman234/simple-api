package main

import (
	"github.com/joho/godotenv"
	"github.com/salamanderman234/simple-api/routes"
)

func main() {
	router := routes.Init()
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	router.Logger.Fatal(router.Start(":3000"))
}
