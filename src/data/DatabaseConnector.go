package data

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDatabaseConnectionString() string {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")
	db := os.Getenv("POSTGRES_DATABASE")

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable gssencmode=allow", host, user, password, db, port)

	return connectionString
}

func ConnectToDatabase() (*gorm.DB, error) {
	connectionString := getDatabaseConnectionString()

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	return db, err
}
