package data

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"mealwhile/errors"
)

func getDatabaseConnectionString() string {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")
	db := os.Getenv("POSTGRES_DATABASE")

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, db, port)

	return connectionString
}

func ConnectToDatabase() (*gorm.DB, error) {
	connectionString := getDatabaseConnectionString()

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		return nil, errors.NewServerError("something went wrong creating the database")
	}

	return db, err
}
