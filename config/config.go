package config

import (
	"fmt"
	"os"
)

func GetDatabaseUrl() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	return url
}
