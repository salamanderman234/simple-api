package main

import (
	"github.com/salamanderman234/simple-api/routes"
)

func main() {
	router := routes.Init()

	router.Logger.Fatal(router.Start(":3000"))
}
