package main

import (
	"fmt"
	"mealwhile/api/controller"
	"mealwhile/data"
	"mealwhile/logic/operations"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

func main() {
	// --------------------
	// Initiate the logger
	// --------------------
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{ForceColors: true})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// --------------------
	// Create all dependencies
	// --------------------

	// Create the database connection
	db, err := data.ConnectToDatabase()

	if err != nil {
		log.Panic("The database connection failed")
		return
	}

	// Create the repositories
	unitRepository := data.NewUnitRepository(db)

	// Create the CrudServices
	unitCrudService := operations.NewCrudService(unitRepository)

	// Create the operations
	unitOperations := operations.NewUnitOperations(unitCrudService)

	// Create the controllers
	groceryController := controller.NewGroceryController(unitOperations)

	// --------------------
	// Create and start the webserver
	// --------------------
	e := echo.New()

	controller.RegisterHandlers(e, &groceryController)

	// Get the port on which the server should run
	portEnv := os.Getenv("PORT")
	if portEnv == "" {
		// Set the default port
		portEnv = "8080"
	}

	port, err := strconv.Atoi(portEnv)

	if err != nil {
		e.Logger.Fatal("There is something wrong with the port config.")
	}

	// Start the server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
