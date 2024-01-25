// **cmd:** This folder contains the main entry point of the application.
// The `main.go` file initializes the application and starts the
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/dzemildupljak/simple_hexa/config"
	"github.com/dzemildupljak/simple_hexa/internal/app/application"
	httphdl "github.com/dzemildupljak/simple_hexa/internal/app/ports/inbound/http"
	persistence "github.com/dzemildupljak/simple_hexa/internal/infrastructure/persistence/postgres"
)

func main() {
	config.LoadEnv()
	config.NewLoggerToVolume()
	config.NewNRApplication()

	// Setup the user service, repository, http handler
	userrepository := persistence.NewUserRepository()
	userservice := application.NewUserService(userrepository)
	userhttpHandler := httphdl.NewUserHTTPHandler(userservice)

	// Create a new router
	router := mux.NewRouter()

	// Register user HTTP handlers
	userhttpHandler.RegisterHandlers(router, config.NRapp)

	// Start the server
	port, valid := os.LookupEnv("APP_PORT")
	if !valid {
		port = "8080"
	}

	fmt.Printf("Server is running on PORT:%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
