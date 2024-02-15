// **cmd:** This folder contains the main entry point of the application.
// The `main.go` file initializes the application and starts the
// @title My Hex App
// @version 1
// @description This is a sample server.
// @termsOfService http://example.com/terms/
// @contact name Developer name email developer@example.com
// @host localhost:8080
// @BasePath /api/v1
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/dzemildupljak/simple_hexa/config"
	_ "github.com/dzemildupljak/simple_hexa/docs" // your generated docs
	"github.com/dzemildupljak/simple_hexa/internal/app/application"
	httphdl "github.com/dzemildupljak/simple_hexa/internal/app/ports/inbound/http"
	persistence "github.com/dzemildupljak/simple_hexa/internal/infrastructure/persistence/postgres"
)

func main() {
	initializeConfiguration()
	router := setupRouter()
	configureSwagger(router)
	configureUserHandler(router)
	startServer(router)
}

func initializeConfiguration() {
	config.LoadEnv()
	config.NewLoggerToVolume()
	config.NewNRApplication()
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

func configureSwagger(router *mux.Router) {
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}

func configureUserHandler(router *mux.Router) {
	// Set up the user service, repository, http handler
	repository := persistence.NewUserRepository()
	service := application.NewUserService(repository)
	httpHandler := httphdl.NewUserHTTPHandler(service)
	httpHandler.RegisterHandlers(router, config.NRapp)
}

func startServer(router *mux.Router) {
	port, valid := os.LookupEnv("APP_PORT")
	if !valid {
		port = "8080"
	}

	fmt.Printf("Server is running on PORT:%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
