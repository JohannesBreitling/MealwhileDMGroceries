package main

import (
	"fmt"
	"mealwhile/api/controller"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	// --------------------
	// Create all dependencies
	// --------------------
	testController := controller.NewTestController()

	// --------------------
	// Create and start the webserver
	// --------------------
	e := echo.New()

	controller.RegisterHandlers(e, &testController)

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
