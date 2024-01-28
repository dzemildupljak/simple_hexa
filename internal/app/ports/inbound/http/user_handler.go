// **inbound:** Inbound ports define interfaces for communication with the external world (e.g., API handlers)
package httphdl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/dzemildupljak/simple_hexa/config"
	"github.com/dzemildupljak/simple_hexa/internal/app/application"
	"github.com/dzemildupljak/simple_hexa/internal/app/domain"
)

// UserHTTPHandler contains HTTP handlers for the application.
type UserHTTPHandler struct {
	userService application.UserService
}

// NewUserHTTPHandler creates a new HTTPHandler with the given UserService.
func NewUserHTTPHandler(userService application.UserService) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: userService,
	}
}

// RegisterHandlers registers HTTP handlers with the provided router.
func (h *UserHTTPHandler) RegisterHandlers(router *mux.Router, nrapp *newrelic.Application) {
	router.HandleFunc(
		newrelic.WrapHandleFunc(
			nrapp,
			"/users",
			config.NrHttpLogger(
				h.GetAllUsersHandler,
			),
		),
	).Methods("GET")

	router.HandleFunc(
		newrelic.WrapHandleFunc(
			nrapp,
			"/users",
			config.NrHttpLogger(
				h.CreateUserHandler,
			),
		),
	).Methods("POST")
	router.HandleFunc(
		newrelic.WrapHandleFunc(
			nrapp,
			"/users/{id}",
			config.NrHttpLogger(
				h.GetUserByIdHandler,
			),
		),
	).Methods("GET")

	router.HandleFunc(
		newrelic.WrapHandleFunc(
			nrapp,
			"/users/email/{email}",
			config.NrHttpLogger(
				h.GetUserByEmailHandler,
			),
		),
	).Methods("GET")
}

func (h *UserHTTPHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	u := &domain.User{}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(u)
	if err != nil {
		http.Error(w, "Error decoding user", http.StatusBadRequest)
		return
	}

	err = h.userService.CreateUser(ctx, u)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User created successfully!\n")
}

func (h *UserHTTPHandler) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	userID := vars["id"]
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	u, err := h.userService.GetUserById(ctx, id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	userJSON, err := json.Marshal(u)
	if err != nil {
		http.Error(w, "Error marshaling user data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}

func (h *UserHTTPHandler) GetUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	uemail := vars["email"]

	u, err := h.userService.GetUserByEmail(ctx, uemail)
	if err != nil {
		http.Error(w, "Error geting user", http.StatusBadRequest)
		return
	}

	uJson, err := json.Marshal(u)
	if err != nil {
		http.Error(w, "Error marshaling user data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(uJson)
}

func (h *UserHTTPHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	txn := newrelic.FromContext(r.Context())

	u, err := h.userService.GetAllUsers(ctx)
	if err != nil {
		txn.NoticeError(err)
		http.Error(w, "Error geting users", http.StatusBadRequest)
		return
	}

	uJson, err := json.Marshal(u)
	if err != nil {
		txn.NoticeError(err)
		http.Error(w, "Error marshaling users data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(uJson)
}
