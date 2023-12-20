// **inbound:** Inbound ports define interfaces for communication with the external world (e.g., API handlers)
package hdlhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dzemildupljak/simple_hexa/internal/app/application"
	"github.com/dzemildupljak/simple_hexa/internal/app/domain"
	"github.com/gorilla/mux"
)

// HTTPHandler contains HTTP handlers for the application.
type HTTPHandler struct {
	userService application.UserService
}

// NewHTTPHandler creates a new HTTPHandler with the given UserService.
func NewHTTPHandler(userService application.UserService) *HTTPHandler {
	return &HTTPHandler{
		userService: userService,
	}
}

// RegisterHandlers registers HTTP handlers with the provided router.
func (h *HTTPHandler) RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/users", h.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", h.GetUserByIDHandler).Methods("GET")
	router.HandleFunc("/users/email/{email}", h.GetUserByEmailHandler).Methods("GET")
	router.HandleFunc("/users", h.GetAllUsersHandler).Methods("GET")
}

func (h *HTTPHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	u := &domain.User{}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(u)
	if err != nil {
		http.Error(w, "Error decoding user", http.StatusBadRequest)
		return
	}

	err = h.userService.CreateUser(u)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User created successfully!\n")
}

func (h *HTTPHandler) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Error geting user1", http.StatusBadRequest)
		return
	}

	u, err := h.userService.GetUserByID(id)
	if err != nil {
		http.Error(w, "Error geting user2", http.StatusBadRequest)
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

func (h *HTTPHandler) GetUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uemail := vars["email"]

	u, err := h.userService.GetUserByEmail(uemail)
	if err != nil {
		http.Error(w, "Error geting user", http.StatusBadRequest)
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

func (h *HTTPHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation similar to the previous example
	fmt.Fprint(w, "Hello, World! GetAllUsersHandler DEPLOYED FROM GH\n")
}
