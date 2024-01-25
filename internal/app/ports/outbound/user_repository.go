// **outbound:** Outbound ports define interfaces for external dependencies (e.g., database interfaces, APIs interfaces)
package outbound

import (
	"errors"

	"github.com/dzemildupljak/simple_hexa/internal/app/domain"
)

var ErrUserNotFound = errors.New("user not found")
var ErrInvalidOperation = errors.New("error invalid operation")

// UserRepository defines the interface for interacting with user data.
type UserRepository interface {
	SaveUser(user *domain.User) error
	GetUserById(userId int) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
}
