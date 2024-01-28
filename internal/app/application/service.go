// **application:** Contains interfaces and implementations for application services.
package application

import (
	"context"

	"github.com/dzemildupljak/simple_hexa/internal/app/domain"
)

// UserService defines the interface for user-related application services.
type UserService interface {
	CreateUser(ctx context.Context, newUser *domain.User) error
	GetUserById(ctx context.Context, userId int) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetAllUsers(ctx context.Context) ([]*domain.User, error)
}
