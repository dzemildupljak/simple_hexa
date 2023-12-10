// **cmd:** This folder contains the main entry point of the application.
// The `main.go` file initializes the application and starts the
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dzemildupljak/simple_hexa/internal/app/application"
	hdlhttp "github.com/dzemildupljak/simple_hexa/internal/app/ports/inbound/http"
	persistence "github.com/dzemildupljak/simple_hexa/internal/infrastructure/persistence/postgres"
	"github.com/gorilla/mux"
)

func main() {
	// Setup the user service and repository
	userRepository := persistence.NewUserRepository()
	userService := application.NewUserService(userRepository)

	// Create a new HTTP handler with the UserService
	httpHandler := hdlhttp.NewHTTPHandler(userService)

	// Create a new router
	router := mux.NewRouter()

	// Register HTTP handlers
	httpHandler.RegisterHandlers(router)

	// Start the server
	port := 80
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
