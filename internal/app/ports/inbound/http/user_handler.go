// Package httphdl **inbound:** Inbound ports define interfaces for communication with the external world (e.g., API handlers)
package httphdl

import (
	"fmt"
	"github.com/dzemildupljak/simple_hexa/config"
	"github.com/dzemildupljak/simple_hexa/internal/app/application"
	"github.com/dzemildupljak/simple_hexa/internal/app/domain"
	httpdto "github.com/dzemildupljak/simple_hexa/internal/app/ports/inbound/http/dto"
	"github.com/dzemildupljak/simple_hexa/utils"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	"net/http"
	"strconv"
	"strings"
)

// UserHttpHandler contains HTTP handlers for the application.
type UserHttpHandler struct {
	userService application.UserService
}

// NewUserHTTPHandler creates a new HTTPHandler with the given UserService.
func NewUserHTTPHandler(userService application.UserService) *UserHttpHandler {
	return &UserHttpHandler{
		userService: userService,
	}
}

func (h *UserHttpHandler) RegisterHandlers(router *mux.Router, nrApp *newrelic.Application) {
	h.applyMiddlewares(router, nrApp)

	// Create a sub router for api/v1/users
	usersRouter := setupUserVersionedRouter(router, "v1")
	h.setupUserRoutes(usersRouter)
}

func setupUserVersionedRouter(router *mux.Router, version string) *mux.Router {
	return router.PathPrefix(fmt.Sprintf("/api/%s/users", version)).Subrouter()
}

func (h *UserHttpHandler) applyMiddlewares(router *mux.Router, nrApp *newrelic.Application) {
	router.Use(
		config.NrHttpLogMiddleware,
		config.NrHttpContextTransaction(nrApp),
	)
}

func (h *UserHttpHandler) setupUserRoutes(usersRouter *mux.Router) {
	usersRouter.HandleFunc("", h.GetAllUsersHandler).Methods("GET")
	usersRouter.HandleFunc("", h.CreateUserHandler).Methods("POST")
	usersRouter.HandleFunc("/{identifier}", h.GetUserHandler).Methods("GET")
}

// CreateUserHandler Create new user
// @Summary Create new user
// @Description Adds a new user to the system
// @Tags user
// @Accept json
// @Produce json
// @Param user body httpdto.CreateUserRequest required "User data"
// @Success 201
// @Failure 400 {string} string "BadRequest: Error creating user"
// @Router /users [post]
func (h *UserHttpHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	u, err := utils.JsonDecode[*httpdto.CreateUserRequest](r)
	if err != nil {
		http.Error(w, "Error decoding user", http.StatusBadRequest)
		return
	}

	usr := domain.NewUser(
		u.Username,
		u.Email,
	)

	err = h.userService.CreateUser(ctx, usr)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetUserHandler handles the request to retrieve user
// @Summary Get user by ID or Email
// @Description Get details of a user by user ID or email address.
// @Tags user
// @Accept  json
// @Produce  json
// @Param identifier path string true "User ID or Email"
// @Success 200 {object} domain.User
// @Failure 400 {string} string "BadRequest: User not found"
// @Router /users/{identifier} [get]
func (h *UserHttpHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	identifier := vars["identifier"] // 'identifier' can be either an ID or an email

	var u *domain.User
	var err error

	// Attempt to parse the identifier as an integer (ID)
	if id, idErr := strconv.Atoi(identifier); idErr == nil {
		// If successful, it's an ID
		u, err = h.userService.GetUserById(ctx, id)
	} else if strings.Contains(identifier, "@") {
		// If the identifier contains '@', treat it as an email
		u, err = h.userService.GetUserByEmail(ctx, identifier)
	} else {
		http.Error(w, "User not found", http.StatusBadRequest)
	}

	// Handle any errors from GetUserById or GetUserByEmail
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	// Send the user data as JSON
	err = utils.JsonEncode(w, http.StatusOK, u)
	if err != nil {
		http.Error(w, "Error encoding", http.StatusBadRequest)
	}
}

// GetAllUsersHandler handles the request to retrieve all users.
// @Summary Get all users
// @Description Retrieves a list of all users in the system.
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} domain.User
// @Failure 400 {string} string "BadRequest: Error getting users"
// @Router /users [get]
func (h *UserHttpHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	u, err := h.userService.GetAllUsers(ctx)
	if err != nil {
		http.Error(w, "Error getting users", http.StatusBadRequest)
		return
	}

	err = utils.JsonEncode(w, http.StatusOK, u)
	if err != nil {
		return
	}
}
