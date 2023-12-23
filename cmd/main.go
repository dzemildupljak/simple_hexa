// **cmd:** This folder contains the main entry point of the application.
// The `main.go` file initializes the application and starts the
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/dzemildupljak/simple_hexa/internal/app/application"
	hdlhttp "github.com/dzemildupljak/simple_hexa/internal/app/ports/inbound/http"
	persistence "github.com/dzemildupljak/simple_hexa/internal/infrastructure/persistence/postgres"
	"github.com/gorilla/mux"
)

func main() {
	nrapp, err := newrelic.NewApplication(
		newrelic.ConfigAppName("myhexaapp"),
		newrelic.ConfigLicense("eu01xxa216e3ba49cbaa44bb5756cf0bFFFFNRAL"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		os.Exit(1)
		return
	}

	// Setup the user service and repository
	userRepository := persistence.NewUserRepository()
	userService := application.NewUserService(userRepository)

	// Create a new HTTP handler with the UserService
	httpHandler := hdlhttp.NewHTTPHandler(userService)

	// Create a new router
	router := mux.NewRouter()

	// Register HTTP handlers
	httpHandler.RegisterHandlers(nrapp, router)

	// Start the server
	port, valid := os.LookupEnv("APP_PORT")
	if !valid {
		port = "8080"
	}
	fmt.Printf("Server is running on PORT:%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
