// **domain:** Defines business domain models and repository interfaces.
package domain

import (
	"errors"
)

var ErrUserNotFound = errors.New("user not found")
var ErrInvalidOperation = errors.New("error invalid operation")

// UserRepository defines the interface for interacting with user data.
type UserRepository interface {
	SaveUser(user *User) error
	GetUserByID(userID int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetAllUsers() ([]*User, error)
}
