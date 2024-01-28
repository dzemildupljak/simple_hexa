// **outbound:** Outbound ports define interfaces for external dependencies (e.g., database interfaces, APIs interfaces)
package outbound

import (
	"context"
	"errors"

	"github.com/dzemildupljak/simple_hexa/internal/app/domain"
)

var ErrUserNotFound = errors.New("erro: user not found")
var ErrInvalidOperation = errors.New("error: invalid operation")
var ErrUniqueFieldConstraint = errors.New("error: one of the fields violates a unique constraint")

// UserRepository defines the interface for interacting with user data.
type UserRepository interface {
	SaveUser(ctx context.Context, user *domain.User) error
	GetUserById(ctx context.Context, userId int) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetAllUsers(ctx context.Context) ([]*domain.User, error)
}
