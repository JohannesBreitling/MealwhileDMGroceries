package main

import (
	"fmt"
	"mealwhile/api/controller"
	"mealwhile/data"
	"mealwhile/logic/operations"
	"net/http"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	flagRepository := data.NewFlagRepository(db)
	groceryRepository := data.NewGroceryRepository(db)

	// Create the CrudServices
	unitCrudService := operations.NewCrudService(unitRepository)
	flagCrudService := operations.NewCrudService(flagRepository)
	groceryCrudService := operations.NewCrudService(groceryRepository)

	// Create the operations
	unitOperations := operations.NewUnitOperations(unitCrudService)
	flagOperations := operations.NewFlagOperations(flagCrudService)
	groceryOperations := operations.NewGroceryOperations(groceryCrudService)
	flagOperations.SetGroceryOperations(&groceryOperations)
	groceryOperations.SetFlagOperations(&flagOperations)

	// Create the controllers
	unitExpectedCrudArguments := controller.ExpectedCrudArguments{
		Create: []string{"name", "abbreviation"},
		Update: []string{"id", "name", "abbreviation"},
	}

	flagExpectedCrudArguments := controller.ExpectedCrudArguments{
		Create: []string{"name", "description"},
		Update: []string{"id", "name", "description"},
	}

	groceryExpectedCrudArguments := controller.ExpectedCrudArguments{
		Create: []string{"name", "flagIds"},
		Update: []string{"id", "name", "flagIds"},
	}

	// TODO recipeExpectedCrudArguments := controller.ExpectedCrudArguments{
	//	Create: []string{"name", "description"},
	//	Update: []string{"id", "name", "description"},
	//}

	unitCrudController := controller.NewCrudController(unitOperations, unitExpectedCrudArguments)
	flagCrudController := controller.NewCrudController(flagOperations, flagExpectedCrudArguments)
	groceryCrudController := controller.NewCrudController(groceryOperations, groceryExpectedCrudArguments)
	// TODO recipeCrudController := controller.NewCrudController()

	groceryController := controller.NewGroceryController(unitOperations, unitCrudController, flagOperations, flagCrudController, groceryOperations, groceryCrudController)

	// --------------------
	// Create and start the webserver
	// --------------------
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

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
