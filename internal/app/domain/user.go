// Package domain **domain:** Defines business domain models and repository interfaces.
package domain

// User represents a user of the application.
// @Description User is the model representing a user in the system.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// NewUser creates a new User instance.
func NewUser(username, email string) *User {
	return &User{
		Username: username,
		Email:    email,
	}
}
