// **cmd:** This folder contains the main entry point of the application.
// The `main.go` file initializes the application and starts the
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dzemildupljak/simple_hexa/config"
	"github.com/dzemildupljak/simple_hexa/internal/app/application"
	hdlhttp "github.com/dzemildupljak/simple_hexa/internal/app/ports/inbound/http"
	persistence "github.com/dzemildupljak/simple_hexa/internal/infrastructure/persistence/postgres"
	"github.com/gorilla/mux"
	"golang.org/x/net/http2"
)

func main() {

	config.NewNRApplication()

	// Setup the user service and repository
	userRepository := persistence.NewUserRepository()
	userService := application.NewUserService(userRepository)

	// Create a new HTTP handler with the UserService
	httpHandler := hdlhttp.NewHTTPHandler(userService)

	// Create a new router
	router := mux.NewRouter()

	// Register HTTP handlers
	httpHandler.RegisterHandlers(config.NRapp, router)

	// Start the server
	port, valid := os.LookupEnv("APP_PORT")
	if !valid {
		port = "8080"
	}

	server := &http.Server{Addr: port}
	http2.ConfigureServer(server, &http2.Server{
		MaxConcurrentStreams: 20,
	})
	fmt.Printf("Server is running on PORT:%s\n", port)
	log.Fatal(server.ListenAndServe())
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
